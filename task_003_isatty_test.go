package isatty_test

import "testing"

// TestTask003 isolates the 句柄文本解析 regression.
func TestTask003(t *testing.T) {
	TestCygwinPipeName(t)
}
