//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty009(t *testing.T) {
	if isCygwinPipeName(`\msys-abc-console0-from-master`) {
		t.Fatal("non-pty token accepted")
	}
}

func TestTaskIsatty009Boundary(t *testing.T) { TestTaskIsatty009(t) }
