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
	path := goconf.VarStringOrDefault("log/", "log", "path")
	backup := goconf.VarBoolOrDefault(false, "log", "backup")
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
	if backup {
		c.OutputPaths = append(c.OutputPaths, path+"info.log")
		c.ErrorOutputPaths = append(c.ErrorOutputPaths, path+"error.log")
	}
	l, err := c.Build()
	if err != nil {
		panic(err)
	}

	logger = l
	
}

func GetLogger() *zap.Logger {
	return logger
}

func Info(v any) {
	Infof("%v", v)
}

func Infof(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logger.Info(s)
}


func Error(v any) {
	Errorf("%v", v)
}

func Errorf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logger.Error(s)
}

func Warn(v any) {
	Warnf("%v", v)
}

func Warnf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logger.Warn(s)
}

func Debug(v any) {
	Debugf("%v", v)
}

func Debugf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logger.Debug(s)
}

func DPanic(v any) {
	DPanicf("%v", v)
}

func DPanicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logger.DPanic(s)
}

func Panic(v any) {
	Panicf("%v", v)
}

func Panicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logger.Panic(s)
}

func Fatal(v any) {
	Fatalf("%v", v)
}

func Fatalf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	logger.Fatal(s)
}
