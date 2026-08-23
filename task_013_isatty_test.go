package isatty_test

import "testing"

// TestTask013 isolates the Windows设备路径 regression.
func TestTask013(t *testing.T) {
	TestTerminal(t)
}

// TestTask013Repeat guards deterministic behavior across repeated calls.
func TestTask013Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask013(t)
	}
}
