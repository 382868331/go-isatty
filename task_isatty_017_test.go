//go:build solaris && !appengine

package isatty

import "testing"

func TestTaskIsatty017(t *testing.T) { _ = IsTerminal(^uintptr(0)) }
