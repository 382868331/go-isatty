package isatty_test

import "testing"

// TestTask013 isolates the Windows设备路径 regression.
func TestTask013(t *testing.T) {
	TestTerminal(t)
}
