package client

import (
	"crypto/tls"
	"net"
	"net/http"
	"runtime"
	"time"
)

func createTransport(timeout, keepalive time.Duration, insecureSkipVerify bool) *http.Transport {
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	if keepalive == 0 {
		keepalive = 30 * time.Second
	}

	dialer := &net.Dialer{
		Timeout:   timeout,
		KeepAlive: keepalive,
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	transport.Proxy = nil

	if insecureSkipVerify {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	return transport
}
