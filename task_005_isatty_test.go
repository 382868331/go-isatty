package isatty_test

import "testing"

// TestTask005 isolates the 探测收尾 regression.
func TestTask005(t *testing.T) {
	TestCygwinPipeName(t)
}
