package isatty_test

import "testing"

// TestTask002 isolates the 描述符集合 regression.
func TestTask002(t *testing.T) {
	TestCygwinPipeName(t)
}

// TestTask002Repeat guards deterministic behavior across repeated calls.
func TestTask002Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask002(t)
	}
}
