package log

type Logger interface {
	Debug(format string, args ...any)

	Info(format string, args ...any)

	Warn(format string, args ...any)

	Error(format string, args ...any)

	Fatal(format string, args ...any)

	Panic(format string, atgs ...any)

	Close() error
}

func New() (Logger, error) {
	return newZap()
}
