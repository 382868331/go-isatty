//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty004(t *testing.T) {
	if !isCygwinPipeName(`\cygwin-abc-pty0-to-master`) {
		t.Fatal("valid cygwin pipe rejected")
	}
}
