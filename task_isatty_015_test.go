//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty015(t *testing.T) {
	if isCygwinPipeName(`\msys-abc-pty0-from-master-nat-`) {
		t.Fatal("later empty suffix ignored")
	}
}
