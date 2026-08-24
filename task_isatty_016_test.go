//go:build (darwin || dragonfly || freebsd || netbsd || openbsd) && !appengine

package isatty

import "testing"

func TestTaskIsatty016(t *testing.T) { _ = IsTerminal(^uintptr(0)) }
