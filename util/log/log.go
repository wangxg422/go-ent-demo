package log

// Debug DebugLevel, Usually only enabled when debugging. Very verbose logging.
func Debug(args ...any) {
	logger.Debug(args...)
}

// Info InfoLevel, General operational entries about what's going on inside the application.
func Info(args ...any) {
	logger.Info(args...)
}

// Warn WarnLevel level. Non-critical entries that deserve eyes.
func Warn(args ...any) {
	logger.Warn(args...)
}

// Error ErrorLevel level. Logs. Used for errors that should definitely be noted.
// Commonly used for hooks to send errors to an error tracking service.
func Error(args ...any) {
	logger.Error(args...)
}

// Fatal FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
// logging level is set to Panic.
func Fatal(args ...any) {
	logger.Fatal(args...)
}

// Debugf DebugLevel, Usually only enabled when debugging. Very verbose logging.
func Debugf(format string, args ...any) {
	logger.Debugf(format, args...)
}

// Infof InfoLevel, General operational entries about what's going on inside the application.
func Infof(format string, args ...any) {
	logger.Infof(format, args...)
}

// Warnf WarnLevel level. Non-critical entries that deserve eyes.
func Warnf(format string, args ...any) {
	logger.Warnf(format, args...)
}

// Errorf ErrorLevel level. Logs. Used for errors that should definitely be noted.
// Commonly used for hooks to send errors to an error tracking service.
func Errorf(format string, args ...any) {
	logger.Errorf(format, args...)
}

// Fatalf FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
// logging level is set to Panic.
func Fatalf(format string, args ...any) {
	logger.Fatalf(format, args...)
}
