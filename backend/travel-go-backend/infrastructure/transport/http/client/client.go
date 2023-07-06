package client

import (
	"context"
	"fmt"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"io"
	"net/http"
	"net/http/cookiejar"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
	"go.uber.org/zap"
	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/net/publicsuffix"
)

type HTTPClient struct {
	httpClient         *http.Client
	timeout            time.Duration
	keepalive          time.Duration
	retryCount         int
	retryAfter         time.Duration
	retryCondition     func(r *http.Response, e error) bool
	insecureSkipVerify bool
}

type ClientOption func(c *HTTPClient)

func NewHTTPClient(opts ...ClientOption) *HTTPClient {
	c := &HTTPClient{
		timeout:    30 * time.Second,
		keepalive:  60 * time.Second,
		retryCount: 2,
		retryAfter: 2 * time.Second,
		retryCondition: func(r *http.Response, e error) bool {
			return r.StatusCode == http.StatusTooManyRequests ||
				r.StatusCode == http.StatusInternalServerError ||
				r.StatusCode == http.StatusBadGateway ||
				r.StatusCode == http.StatusFound ||
				r.StatusCode == http.StatusServiceUnavailable ||
				r.StatusCode == http.StatusGatewayTimeout
		},
		insecureSkipVerify: true, // Disable this optional in future if going to the production to secure
	}

	for _, opt := range opts {
		opt(c)
	}

	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	c.httpClient = apmhttp.WrapClient(&http.Client{
		Transport: createTransport(c.timeout, c.timeout, c.insecureSkipVerify),
		Jar:       cookieJar,
	})

	return c
}

func SetTimeout(t time.Duration) ClientOption {
	return func(c *HTTPClient) {
		c.timeout = t
	}
}

func SetKeepAlive(t time.Duration) ClientOption {
	return func(c *HTTPClient) {
		c.keepalive = t
	}
}

func SetRetry(count int) ClientOption {
	return func(c *HTTPClient) {
		c.retryCount = count
	}
}

func SetRetryAfter(t time.Duration) ClientOption {
	return func(c *HTTPClient) {
		c.retryAfter = t
	}
}

func SetRetryCondition(fn func(r *http.Response, e error) bool) ClientOption {
	return func(c *HTTPClient) {
		c.retryCondition = fn
	}
}

func SetInsecureSkipVerify() ClientOption {
	return func(c *HTTPClient) {
		c.insecureSkipVerify = true
	}
}

func (c *HTTPClient) execute(ctx context.Context, r *http.Request) (*http.Response, error) {
	return ctxhttp.Do(ctx, c.httpClient, r)
}

func (c *HTTPClient) retry(ctx context.Context, r *http.Request) (*http.Response, error) {
	traceLogger := logger.TraceLogger(ctx)

	var (
		retried int32 = 0
		res     *http.Response
		err     error
	)

	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("context deadline")
		default:
			res, err = c.execute(ctx, r)
			// Return result if has response, and no error
			if res != nil && err == nil {
				return res, err
			}

			// Return err if has no res
			if err != nil && res == nil {
				return nil, err
			}

			//If has both res and error, do retry
			traceLogger.Warn(fmt.Sprintf("Retry reconnect to HTTP Address: %s", r.RequestURI))

			atomic.AddInt32(&retried, 1)
			if !c.retryCondition(res, err) || retried >= int32(c.retryCount) {
				return res, err
			}
			time.Sleep(c.retryAfter)
		}
	}
}

type ExecuteOption func(req *http.Request, resp *http.Response, e error) error

func (c *HTTPClient) Execute(ctx context.Context, req *http.Request, opts ...ExecuteOption) (*http.Response, error) {
	traceLogger := logger.TraceLogger(ctx)

	span, ctx := apm.StartSpan(ctx, "http.Execute", "http.External")
	defer span.End()

	res, err := c.retry(ctx, req)
	span.Context.SetHTTPRequest(req)
	if res != nil {
		span.Context.SetHTTPStatusCode(res.StatusCode)
	}

	defer func() {
		if res != nil {
			if err := res.Body.Close(); err != nil {
				go apm.CaptureError(ctx, errors.Wrap(err, "call http.Response.Body.Close() has failed")).Send()
			}
		}
	}()

	if err != nil {
		return res, err
	}

	for _, opt := range opts {
		if err := opt(req, res, err); err != nil {
			return res, err
		}
	}

	payload, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	traceLogger.Debug("Debug soap response", zap.String("body", string(payload)), zap.Int("status_code", res.StatusCode))

	return res, nil
}
