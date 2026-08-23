package isatty_test

import "testing"

// TestTask015 isolates the 无效句柄 regression.
func TestTask015(t *testing.T) {
	TestCygwinPipeName(t)
}
