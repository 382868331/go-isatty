package isatty_test

import "testing"

// TestTask012 isolates the Cygwin管道名 regression.
func TestTask012(t *testing.T) {
	TestCygwinPipeName(t)
}

// TestTask012Repeat guards deterministic behavior across repeated calls.
func TestTask012Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask012(t)
	}
}
