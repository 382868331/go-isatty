package isatty_test

import "testing"

// TestTask014 isolates the 终端尺寸 regression.
func TestTask014(t *testing.T) {
	TestCygwinPipeName(t)
}
