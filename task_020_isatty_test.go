package isatty_test

import "testing"

// TestTask020 isolates the 平台能力版本 regression.
func TestTask020(t *testing.T) {
	TestCygwinPipeName(t)
}
