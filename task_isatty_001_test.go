//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty001(t *testing.T) {
	if IsTerminal(^uintptr(0)) {
		t.Fatal("invalid handle classified as terminal")
	}
}

func TestTaskIsatty001Boundary(t *testing.T) { TestTaskIsatty001(t) }
