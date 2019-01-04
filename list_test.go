package main

import (
	"sync"
	"testing"
	"unsafe"
)

type IntKey struct {
	v int
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

func TestInsert(t *testing.T) {
	list := NewList()
	wg := &sync.WaitGroup{}
	n := 1024
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			list.Insert(&IntKey{i % 8})
			wg.Done()
		}(i)
	}
	wg.Wait()
	p := list.head
	c := 0
	for {
		if p == nil {
			break
		}
		p = (*Node)(p).next
		c += 1
	}
	if c != n+2 {
		t.Errorf("c must be n+2. but got %d", c)
	}
}
