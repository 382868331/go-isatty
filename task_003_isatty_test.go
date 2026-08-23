package isatty_test

import "testing"

// TestTask003 isolates the 句柄文本解析 regression.
func TestTask003(t *testing.T) {
	TestCygwinPipeName(t)
}

// TestTask003Repeat guards deterministic behavior across repeated calls.
func TestTask003Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask003(t)
	}
}
