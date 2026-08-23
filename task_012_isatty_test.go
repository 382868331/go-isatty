package isatty_test

import "testing"

// TestTask012 isolates the Cygwin管道名 regression.
func TestTask012(t *testing.T) {
	TestCygwinPipeName(t)
}
