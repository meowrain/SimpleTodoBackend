package loggers

import "testing"

func TestLogger(t *testing.T) {
	TodoLogger.Info("This is a info logger")
	Debug("This is a debug logger")
}
