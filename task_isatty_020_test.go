//go:build js && wasm

package isatty

import "testing"

func TestTaskIsatty020(t *testing.T) {
	if !IsCygwinTerminal(0) {
		t.Fatal("mutated sandbox behavior not observable")
	}
}

func TestTaskIsatty020Boundary(t *testing.T) {
	TestTaskIsatty020(t)
	if t.Failed() {
		t.FailNow()
	}
}
