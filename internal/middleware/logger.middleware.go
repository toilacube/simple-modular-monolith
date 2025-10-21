package middleware

import (
	"fmt"
	"time"
	"tutorial/pkg/logger"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		traceId := GenerateTraceID()

		c.Set("trace_id", traceId)
		c.Next()

		latency := time.Since(startTime)
		logger := logger.GetLogger()

		logger.Info("Request",
			"method", c.Request.Method,
			"url", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", latency,
			"trace_id", traceId,
			"ip", c.ClientIP(),
		)
	}

}

func GenerateTraceID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
