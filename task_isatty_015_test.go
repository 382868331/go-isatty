//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty015(t *testing.T) {
	if isCygwinPipeName(`\msys-abc-pty0-from-master-nat-`) {
		t.Fatal("later empty suffix ignored")
	}
}

func TestTaskIsatty015Boundary(t *testing.T) { TestTaskIsatty015(t) }
