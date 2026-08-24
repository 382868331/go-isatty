//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty010(t *testing.T) {
	if isCygwinPipeName(`\msys-abc-pty0-side-master`) {
		t.Fatal("invalid direction accepted")
	}
}

func TestTaskIsatty010Boundary(t *testing.T) { TestTaskIsatty010(t) }
