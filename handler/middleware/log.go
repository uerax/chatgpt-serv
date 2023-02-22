package middleware

import (
	"bytes"
	"io"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/uerax/chatgpt-prj/logger"
)

func ZapLogger() gin.HandlerFunc {
	return ginzap.GinzapWithConfig(logger.GetLogger(), &ginzap.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
		Context: ginzap.Fn(func(c *gin.Context) []zapcore.Field {
		  fields := []zapcore.Field{}
		  // log request ID
		  if requestID := c.Writer.Header().Get("X-Request-Id"); requestID != "" {
			fields = append(fields, zap.String("request_id", requestID))
		  }
	
		  // log trace and span ID
		  if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
			fields = append(fields, zap.String("trace_id", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()))
			fields = append(fields, zap.String("span_id", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()))
		  }
	
		  // log request body
		  var body []byte
		  var buf bytes.Buffer
		  tee := io.TeeReader(c.Request.Body, &buf)
		  body, _ = io.ReadAll(tee)
		  c.Request.Body = io.NopCloser(&buf)
		  fields = append(fields, zap.String("body", string(body)))
	
		  return fields
		}),
	  })
}

func ZapLoggerRec() gin.HandlerFunc {
	return ginzap.RecoveryWithZap(logger.GetLogger(), true)
}