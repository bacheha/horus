package logger

import (
	"log"

	"go.uber.org/zap"
)

type Logger struct {
	log *zap.SugaredLogger
	Std *zap.Logger
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.log.Infow(message, args...)
}

func (l *Logger) Infof(message string, args ...interface{}) {
	l.log.Infof(message, args...)
}

func (l *Logger) Fatalf(message string, args ...interface{}) {
	l.log.Fatalf(message, args...)
}

func (l *Logger) ToStd() *log.Logger {
	return zap.NewStdLog(l.Std)
}

func New() (*Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	logger := log.Sugar()
	return &Logger{
		Std: log,
		log: logger,
	}, nil
}
