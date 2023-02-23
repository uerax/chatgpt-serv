package middleware

import (
	"errors"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uerax/chatgpt-prj/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LogInit() {
	log = logger.GetLogger()
}

var log *zap.Logger

func ZapLogger() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		end := time.Now()

		fields := []zapcore.Field{
			zap.Int("code", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.Duration("latency", end.Sub(start)),
			zap.String("path", path),
			zap.String("param", raw),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
		}

		fields = append(fields, zap.String("date", start.Format("2006-01-02 15:04:05")))

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				log.Error(e, fields...)
			}
		} else {
			log.Info(path, fields...)
		}
		
	}
}

func defaultHandleRecovery(c *gin.Context, err any) {
	c.AbortWithStatus(http.StatusInternalServerError)
}

func ZapRecovery() gin.HandlerFunc {
	
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				if brokenPipe {
					log.Error(c.Request.URL.Path,
						zap.Any("error", err),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
	
				} else {
					defaultHandleRecovery(c, err)
				}
			}
		}()
		c.Next()
	}
}