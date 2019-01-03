package main

import (
	"unsafe"
)

type Key interface {
	LessThan(key Key) bool
}

type Node struct {
	key  Key
	next unsafe.Pointer
}

type List struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewList() *List {
	tail := unsafe.Pointer(&Node{nil, nil})
	head := unsafe.Pointer(&Node{nil, tail})
	list := &List{head, tail}
	return list
}

func (list *List) findLeftAndRight(key Key) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	t := list.head
	tNext := (*Node)(list.head).next
	var leftNode unsafe.Pointer = nil
	var leftNodeNext unsafe.Pointer = nil
	for {
		if !marked(tNext) {
			leftNode = t
			leftNodeNext = tNext
		}
		t = unmark(tNext)
		if t == list.tail {
			break
		}
		tNext = (*Node)(t).next
		if !marked(tNext) && !(*Node)(t).key.LessThan(key) {
			break
		}
	}
	return leftNode, leftNodeNext, t
}
