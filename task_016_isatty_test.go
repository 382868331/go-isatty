package isatty_test

import "testing"

// TestTask016 isolates the 终端模式恢复 regression.
func TestTask016(t *testing.T) {
	TestTerminal(t)
}
