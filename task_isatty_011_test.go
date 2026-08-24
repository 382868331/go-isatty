//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty011(t *testing.T) {
	if isCygwinPipeName(`\msys-abc-pty0-from-worker`) {
		t.Fatal("invalid role accepted")
	}
}

func TestTaskIsatty011Boundary(t *testing.T) { TestTaskIsatty011(t) }
