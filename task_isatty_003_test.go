//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty003(t *testing.T) {
	if !isCygwinPipeName(`\msys-abc-pty0-from-master`) {
		t.Fatal("valid msys pipe rejected")
	}
}
