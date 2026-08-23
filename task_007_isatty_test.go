package isatty_test

import "testing"

// TestTask007 isolates the 设备名Unicode regression.
func TestTask007(t *testing.T) {
	TestTerminal(t)
}
