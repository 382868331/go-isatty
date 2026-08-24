//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty008(t *testing.T) {
	if isCygwinPipeName(`\msys--pty0-from-master`) {
		t.Fatal("empty identity accepted")
	}
}
