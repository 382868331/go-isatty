//go:build plan9

package isatty

import "testing"

func TestTaskIsatty018(t *testing.T) { _ = IsTerminal(^uintptr(0)) }
