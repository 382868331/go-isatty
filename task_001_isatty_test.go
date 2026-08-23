package isatty_test

import "testing"

// TestTask001 isolates the 空描述符 regression.
func TestTask001(t *testing.T) {
	TestTerminal(t)
}
