package isatty_test

import "testing"

// TestTask011 isolates the 并发探测 regression.
func TestTask011(t *testing.T) {
	TestCygwinPipeName(t)
}
