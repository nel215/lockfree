package list

import (
	"testing"
	"unsafe"
)

func TestMark(t *testing.T) {
	s := "abc"
	p1 := unsafe.Pointer(&s)
	if marked(p1) {
		t.Errorf("p1 must be unmarked")
	}
	p2 := mark(p1)
	if !marked(p2) {
		t.Errorf("p must be marked")
	}
	p3 := unmark(p2)
	if p1 != p3 {
		t.Errorf("p3 must be equal p1")
	}
}
