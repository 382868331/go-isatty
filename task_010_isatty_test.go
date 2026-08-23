package isatty_test

import "testing"

// TestTask010 isolates the 设备遍历 regression.
func TestTask010(t *testing.T) {
	TestTerminal(t)
}

// TestTask010Repeat guards deterministic behavior across repeated calls.
func TestTask010Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask010(t)
	}
}
