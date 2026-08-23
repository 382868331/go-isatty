package isatty_test

import "testing"

// TestTask018 isolates the 探测缓存 regression.
func TestTask018(t *testing.T) {
	TestCygwinPipeName(t)
}
