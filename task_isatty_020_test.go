//go:build js && wasm

package isatty

import "testing"

func TestTaskIsatty020(t *testing.T) {
	if !IsCygwinTerminal(0) {
		t.Fatal("mutated sandbox behavior not observable")
	}
}
