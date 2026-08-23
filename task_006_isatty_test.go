package isatty_test

import "testing"

// TestTask006 isolates the 终端候选去重 regression.
func TestTask006(t *testing.T) {
	TestCygwinPipeName(t)
}
