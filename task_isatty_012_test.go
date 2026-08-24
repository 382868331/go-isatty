//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty012(t *testing.T) {
	if isCygwinPipeName(`\msys-abc-pty0-from-master-`) {
		t.Fatal("empty suffix accepted")
	}
}

func TestTaskIsatty012Boundary(t *testing.T) { TestTaskIsatty012(t) }
