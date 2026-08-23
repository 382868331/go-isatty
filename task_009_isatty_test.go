package isatty_test

import "testing"

// TestTask009 isolates the 输入输出句柄 regression.
func TestTask009(t *testing.T) {
	TestCygwinPipeName(t)
}

// TestTask009Repeat guards deterministic behavior across repeated calls.
func TestTask009Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask009(t)
	}
}
