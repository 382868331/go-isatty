package isatty_test

import "testing"

// TestTask009 isolates the 输入输出句柄 regression.
func TestTask009(t *testing.T) {
	TestCygwinPipeName(t)
}
