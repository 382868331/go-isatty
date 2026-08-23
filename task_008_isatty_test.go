package isatty_test

import "testing"

// TestTask008 isolates the 句柄边界 regression.
func TestTask008(t *testing.T) {
	TestCygwinPipeName(t)
}
