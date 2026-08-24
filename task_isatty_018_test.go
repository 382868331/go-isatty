//go:build plan9

package isatty

import "testing"

func TestTaskIsatty018(t *testing.T) { _ = IsTerminal(^uintptr(0)) }

func TestTaskIsatty018Boundary(t *testing.T) {
	TestTaskIsatty018(t)
	if t.Failed() {
		t.FailNow()
	}
}
