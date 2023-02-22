package logger

import (
	"fmt"

	"github.com/uerax/goconf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init() {

	format := goconf.VarStringOrDefault("2006-01-02 15:04:05", "log", "format")
	// path := goconf.VarStringOrDefault("/etc/", "log", "path") 
	
	lv := goconf.VarStringOrDefault("info", "log", "level")
	lvMap := map[string]zapcore.Level{
		"info":zapcore.InfoLevel,
		"debug":zapcore.DebugLevel,
		"warn":zapcore.WarnLevel,
		"error":zapcore.ErrorLevel,
		"dpanic":zapcore.DPanicLevel,
		"panic":zapcore.PanicLevel,
		"fatal":zapcore.FatalLevel,
	}
	level := zapcore.WarnLevel
	if _, ok := lvMap[lv]; ok {
		level = lvMap[lv]
	}

	c := zap.NewProductionConfig()

	if !goconf.VarBoolOrDefault(false, "log", "json") {
		c.Encoding = "console"
	}

	c.Level = zap.NewAtomicLevelAt(level)
	c.DisableStacktrace = true
	if !goconf.VarBoolOrDefault(false, "dev") {
		c.Development = false
	}
	c.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(format)
	c.EncoderConfig.CallerKey = "file"
	c.EncoderConfig.TimeKey = "date"

	l, err := c.Build()
	if err != nil {
		panic(err)
	}

	logger = l
	
}

func GetLogger() *zap.Logger {
	return logger
}

func Info(v interface{}) {
	s := fmt.Sprintf("%v", v)
	logger.Info(s)
}

func Error(v interface{}) {
	s := fmt.Sprintf("%v", v)
	logger.Error(s)
}