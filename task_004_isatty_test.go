package isatty_test

import "testing"

// TestTask004 isolates the 探测取消 regression.
func TestTask004(t *testing.T) {
	TestTerminal(t)
}
