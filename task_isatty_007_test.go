//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty007(t *testing.T) {
	if isCygwinPipeName(`\random-abc-pty0-from-master`) {
		t.Fatal("unrelated pipe accepted")
	}
}
