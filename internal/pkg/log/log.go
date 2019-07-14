package log

var logger logTemp

type logTemp interface {
	Debug(args []interface{})
	Debugf(format string, args []interface{})
	Info(args []interface{})
	Infof(format string, args []interface{})
	Warn(args []interface{})
	Warnf(format string, args []interface{})
	Error(args []interface{})
	Errorf(format string, args []interface{})
	Panic(args []interface{})
	Panicf(format string, args []interface{})
}

// New create log instance
func New(pkg string) {
	if pkg == "zap" {
		NewZapLogger()
		return
	}
	NewZapLogger()
}

// Close end of logger
func Close(pkg string) {
	if pkg == "zap" {
		ZapLogSync()
	}
}

// Debug log debug level message
func Debug(args ...interface{}) {
	logger.Debug(args)
}

// Debugf log debug level message with templated format
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

// Info log info level message with templated format
func Info(args ...interface{}) {
	logger.Info(args)
}

// Infof log info level message with templated format
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args)
}

// Warn log warn level message
func Warn(args ...interface{}) {
	logger.Warn(args)
}

// Warnf log warn level message with templated format
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args)
}

// Error log Error level message
func Error(args ...interface{}) {
	logger.Error(args)
}

// Errorf log error level message with templated format
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args)
}

// Panic log Error level message
func Panic(args ...interface{}) {
	logger.Panic(args)
}

// Panicf log error level message with templated format
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args)
}
