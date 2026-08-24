//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty013(t *testing.T) {
	if !isCygwinPipeName(`\msys-abc-pty0-from-master-nat`) {
		t.Fatal("valid suffix rejected")
	}
}
