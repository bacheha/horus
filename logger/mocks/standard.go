package mocks

import "log"

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(message string, args ...interface{}) {

}

func (l *Logger) Warn(message string, args ...interface{}) {

}

func (l *Logger) Error(message string, args ...interface{}) {

}

func (l *Logger) Infof(template string, values ...interface{}) {

}

func (l *Logger) Warnf(template string, values ...interface{}) {

}

func (l *Logger) Errorf(template string, values ...interface{}) {

}

func (l *Logger) GetStdLogger() *log.Logger {
	return nil
}
