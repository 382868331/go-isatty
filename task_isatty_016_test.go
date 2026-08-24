//go:build (darwin || dragonfly || freebsd || netbsd || openbsd) && !appengine

package isatty

import "testing"

func TestTaskIsatty016(t *testing.T) { _ = IsTerminal(^uintptr(0)) }

func TestTaskIsatty016Boundary(t *testing.T) {
	TestTaskIsatty016(t)
	if t.Failed() {
		t.FailNow()
	}
}
