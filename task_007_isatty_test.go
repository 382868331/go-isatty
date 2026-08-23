package isatty_test

import "testing"

// TestTask007 isolates the 设备名Unicode regression.
func TestTask007(t *testing.T) {
	TestTerminal(t)
}

// TestTask007Repeat guards deterministic behavior across repeated calls.
func TestTask007Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask007(t)
	}
}
