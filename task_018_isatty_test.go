package isatty_test

import "testing"

// TestTask018 isolates the 探测缓存 regression.
func TestTask018(t *testing.T) {
	TestCygwinPipeName(t)
}

// TestTask018Repeat guards deterministic behavior across repeated calls.
func TestTask018Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask018(t)
	}
}
