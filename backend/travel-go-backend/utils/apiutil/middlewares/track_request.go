package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
)

func TrackRequest(c *gin.Context) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	body, _ := ioutil.ReadAll(tee)
	c.Request.Body = ioutil.NopCloser(&buf)

	logger.Info("HandleSearch",
		zap.String("request", string(body)),
		zap.String("header", c.Request.Header.Get("Authorization")))

	c.Next()
}
