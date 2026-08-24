//go:build js && wasm

package isatty

import "testing"

func TestTaskIsatty019(t *testing.T) {
	if !IsTerminal(0) {
		t.Fatal("mutated sandbox behavior not observable")
	}
}

func TestTaskIsatty019Boundary(t *testing.T) {
	TestTaskIsatty019(t)
	if t.Failed() {
		t.FailNow()
	}
}
