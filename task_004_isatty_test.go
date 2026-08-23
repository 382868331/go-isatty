package isatty_test

import "testing"

// TestTask004 isolates the 探测取消 regression.
func TestTask004(t *testing.T) {
	TestTerminal(t)
}

// TestTask004Repeat guards deterministic behavior across repeated calls.
func TestTask004Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask004(t)
	}
}
