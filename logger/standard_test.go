package logger

import (
	"bytes"
	"testing"
)

func TestStandardLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewStandardLogger()
	logger.log.SetOutput(&buf)

	// write logs
	logger.Info("some log line", "key", "value")
	logger.Warn("some log line", "key", "value")
	logger.Error("some log line", "key", "value")
	logger.Infof("hello %s", "world")
	logger.Warnf("hello %s", "world")
	logger.Errorf("hello %s", "world")

	if buf.String() == "" {
		t.Fatal("logger should not have empty buffer")
	}

	if !isStandardLogger(logger.GetStdLogger()) {
		t.Fatal("logger should be of type log/logger")
	}
}
