package log

import (
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLog struct {
	logger *zap.Logger
}

var zapLogger *zapLog

// NewZapLogger create log instance
func NewZapLogger() {
	if logger != nil {
		return
	}
	opts := []zap.Option{
		zap.AddCallerSkip(2),
	}
	l, err := newProductionConfig().Build(opts...)
	if err != nil {
		log.Fatal(err)
	}
	zapLogger = &zapLog{logger: l}
	logger = zapLogger
}

// ZapLogSync sync
func ZapLogSync() {
	if logger == nil {
		return
	}
	zapLogger.logger.Sync()
}

// newProductionConfig production logging configuration.
// Logging is enabled at DebugLevel and above.
//
// It uses a JSON encoder, writes to standard error.
// Stacktraces are automatically included on logs of ErrorLevel and above.
func newProductionConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    newProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// newProductionEncoderConfig returns an opinionated EncoderConfig for
// production environments.
func newProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// timeEncoder serializes a time.Time to a UTC time with 2006-01-02T15:04:05.000 formatted string
// with millisecond precision.
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000"))
}

func (z *zapLog) Debug(args []interface{}) {
	z.logger.Debug(z.composeLogMsg("", args))
}

func (z *zapLog) Debugf(format string, args []interface{}) {
	z.logger.Debug(z.composeLogMsg(format, args))
}

func (z *zapLog) Info(args []interface{}) {
	z.logger.Info(z.composeLogMsg("", args))
}

func (z *zapLog) Infof(format string, args []interface{}) {
	z.logger.Info(z.composeLogMsg(format, args))
}

func (z *zapLog) Warn(args []interface{}) {
	z.logger.Warn(z.composeLogMsg("", args))
}

func (z *zapLog) Warnf(format string, args []interface{}) {
	z.logger.Warn(z.composeLogMsg(format, args))
}

func (z *zapLog) Error(args []interface{}) {
	z.logger.Error(z.composeLogMsg("", args))
}

func (z *zapLog) Errorf(format string, args []interface{}) {
	z.logger.Error(z.composeLogMsg(format, args))
}

func (z *zapLog) Panic(args []interface{}) {
	z.logger.Panic(z.composeLogMsg("", args))
}

func (z *zapLog) Panicf(format string, args []interface{}) {
	z.logger.Panic(z.composeLogMsg(format, args))
}

func (z *zapLog) composeLogMsg(template string, fmtArgs []interface{}) string {
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}
	return msg
}
