package isatty_test

import "testing"

// TestTask017 isolates the 管道名匹配 regression.
func TestTask017(t *testing.T) {
	TestCygwinPipeName(t)
}
