package logger

type Logger interface {
	// Info log info message
	Info(info string)
	// Error log error message
	Error(err error)
	// ErrorOp log error message with operation name
	ErrorOp(op string, err error)
	// Fatal log error message and exit
	Fatal(err error)
	// FatalOp log error message with operation name and exit
	FatalOp(op string, err error)
}
