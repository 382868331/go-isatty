//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty006(t *testing.T) {
	if !isCygwinPipeName(`\Device\NamedPipe\cygwin-abc-pty0-to-master`) {
		t.Fatal("device cygwin pipe rejected")
	}
}

func TestTaskIsatty006Boundary(t *testing.T) { TestTaskIsatty006(t) }
