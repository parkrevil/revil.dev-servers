package logger

type Logger struct {
	tempVal int
}

func NewLogger() *Logger {
	return &Logger{}
}
