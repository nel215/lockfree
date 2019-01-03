package main

import (
	"testing"
	"unsafe"
)

type IntKey struct {
	v int64
}

func (key *IntKey) LessThan(other Key) bool {
	return key.v < other.(*IntKey).v
}

func TestFindLeftAndRightWithDeletedNode(t *testing.T) {
	list := NewList()
	node := unsafe.Pointer(&Node{&IntKey{2}, list.tail})
	(*Node)(list.head).next = node
	(*Node)(node).next = mark(list.tail)
	leftNode, leftNodeNext, rightNode := list.findLeftAndRight(&IntKey{1})
	if leftNode != list.head {
		t.Errorf("leftNode must be head")
	}
	if leftNodeNext != (*Node)(list.head).next {
		t.Errorf("leftNodeNext must be middle node")
	}
	if rightNode != list.tail {
		t.Errorf("rightNode must be tail")
	}
}

func TestFindLeftAndRightWith(t *testing.T) {
	list := NewList()
	node := unsafe.Pointer(&Node{&IntKey{2}, list.tail})
	(*Node)(list.head).next = node
	leftNode, leftNodeNext, rightNode := list.findLeftAndRight(&IntKey{1})
	if leftNode != list.head {
		t.Errorf("leftNode must be head")
	}
	if leftNodeNext != (*Node)(list.head).next {
		t.Errorf("leftNodeNext must be middle node")
	}
	if rightNode != (*Node)(list.head).next {
		t.Errorf("rightNode must be middle node")
	}
}
