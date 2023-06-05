package logger

import "testing"

func TestLoggerConfig(t *testing.T) {
	log := NewLogger()
	log.Debug("Testing debug")
	log.Info("Testing info")
	log.Warn("Testing Warn")
	log.Fatal("Testing Fatal")
}
