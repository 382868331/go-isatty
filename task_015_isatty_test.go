package isatty_test

import "testing"

// TestTask015 isolates the 无效句柄 regression.
func TestTask015(t *testing.T) {
	TestCygwinPipeName(t)
}

// TestTask015Repeat guards deterministic behavior across repeated calls.
func TestTask015Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask015(t)
	}
}
