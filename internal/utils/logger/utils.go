package logger

func GetLogger() Logger {
	return log
}

func Debug(args ...any) {
	log.Debug(args...)
}

func Debugf(format string, args ...any) {
	log.Debugf(format, args...)
}

func Info(args ...any) {
	log.Info(args...)
}

func Infof(format string, args ...any) {
	log.Infof(format, args...)
}

func Warn(args ...any) {
	log.Warn(args...)
}

func Warnf(format string, args ...any) {
	log.Warnf(format, args...)
}

func Error(args ...any) {
	log.Error(args...)
}

func Errorf(format string, args ...any) {
	log.Errorf(format, args...)
}

func Fatal(args ...any) {
	log.Fatal(args...)
}

func Fatalln(args ...any) {
	log.Fatalln(args...)
}

func Fatalf(format string, args ...any) {
	log.Fatalf(format, args...)
}

func Sync() error {
	return log.Sync()
}
