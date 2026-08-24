//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty005(t *testing.T) {
	if !isCygwinPipeName(`\Device\NamedPipe\msys-abc-pty0-from-master`) {
		t.Fatal("device msys pipe rejected")
	}
}

func TestTaskIsatty005Boundary(t *testing.T) { TestTaskIsatty005(t) }
