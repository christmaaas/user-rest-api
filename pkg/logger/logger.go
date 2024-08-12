package logger

type Logger interface {
	Debug(...any)
	Debugf(string, ...any)

	Info(...any)
	Infof(string, ...any)

	Warning(...any)
	Warningf(string, ...any)

	Error(...any)
	Errorf(string, ...any)

	Fatal(...any)
	Fatalf(string, ...any)

	Panic(...any)
	Panicf(string, ...any)

	Trace(...any)
	Tracef(string, ...any)
}
