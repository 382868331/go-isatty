//go:build windows && !appengine

package isatty

import "testing"

func TestTaskIsatty002(t *testing.T) {
	saved := procNtQueryObject
	procNtQueryObject = nil
	defer func() { procNtQueryObject = saved }()
	if _, err := getFileNameByHandle(0); err == nil {
		t.Fatal("unsupported query returned nil error")
	}
}

func TestTaskIsatty002Boundary(t *testing.T) { TestTaskIsatty002(t) }
