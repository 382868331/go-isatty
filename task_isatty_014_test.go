//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty014(t *testing.T) {
	if !isCygwinPipeName(`\msys-abc-pty0-from-master`) {
		t.Fatal("canonical five-token pipe rejected")
	}
}
