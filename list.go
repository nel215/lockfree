package main

import (
	"sync/atomic"
	"unsafe"
)

type Key interface {
	LessThan(key Key) bool
	Equal(key Key) bool
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

func (list *List) search(searchKey Key) (unsafe.Pointer, unsafe.Pointer) {
	// return two references left node and right node.
	// left.Key < key and key <= right.Key.
	// both nodes must be unmarked.
	// the right node must be the immediate successor of the left node
	for {
		leftNode, leftNodeNext, rightNode := list.findLeftAndRight(searchKey)
		// check nodes are adjacent
		if leftNodeNext == rightNode {
			if (rightNode != list.tail) && marked((*Node)(rightNode).next) {
				continue
			} else {
				return leftNode, rightNode
			}
		}
		// remove marked nodes
		if atomic.CompareAndSwapPointer(&(*Node)(leftNode).next, leftNodeNext, rightNode) {
			if (rightNode != list.tail) && marked((*Node)(rightNode).next) {
				continue
			} else {
				return leftNode, rightNode
			}
		}
	}
}

func (list *List) Insert(key Key) {
	newNode := unsafe.Pointer(&Node{key, nil})

	for {
		leftNode, rightNode := list.search(key)
		(*Node)(newNode).next = rightNode
		if atomic.CompareAndSwapPointer(&(*Node)(leftNode).next, rightNode, newNode) {
			return
		}
	}
}

func (list *List) Delete(key Key) bool {
	for {
		leftNode, rightNode := list.search(key)
		if rightNode == list.tail || !(*Node)(rightNode).key.Equal(key) {
			return false
		}
		rightNodeNext := (*Node)(rightNode).next
		if marked(rightNodeNext) {
			continue
		}
		if atomic.CompareAndSwapPointer(&(*Node)(rightNode).next, rightNodeNext, mark(rightNodeNext)) {
			if !atomic.CompareAndSwapPointer(&(*Node)(leftNode).next, rightNode, rightNodeNext) {
				list.search(key)
			}
			return true
		}
	}
}
