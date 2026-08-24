//go:build solaris && !appengine

package isatty

import "testing"

func TestTaskIsatty017(t *testing.T) { _ = IsTerminal(^uintptr(0)) }

func TestTaskIsatty017Boundary(t *testing.T) {
	TestTaskIsatty017(t)
	if t.Failed() {
		t.FailNow()
	}
}
