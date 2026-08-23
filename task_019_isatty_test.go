package isatty_test

import "testing"

// TestTask019 isolates the 空设备集 regression.
func TestTask019(t *testing.T) {
	TestTerminal(t)
}
