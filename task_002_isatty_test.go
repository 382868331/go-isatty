package isatty_test

import "testing"

// TestTask002 isolates the 描述符集合 regression.
func TestTask002(t *testing.T) {
	TestCygwinPipeName(t)
}
