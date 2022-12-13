package logger

import (
	"log"

	"go.uber.org/zap"
)

type Logger struct {
	sugar *zap.SugaredLogger
	log   *zap.Logger
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.sugar.Infow(message, args...)
}

func (l *Logger) Infof(message string, args ...interface{}) {
	l.sugar.Infof(message, args...)
}

func (l *Logger) Fatalf(message string, args ...interface{}) {
	l.sugar.Fatalf(message, args...)
}

func (l *Logger) GetLogger() *zap.Logger {
	return l.log
}

func (l *Logger) GetStdLogger() *log.Logger {
	return zap.NewStdLog(l.log)
}

func New() (*Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	logger := log.Sugar()
	return &Logger{sugar: logger, log: log}, nil
}
