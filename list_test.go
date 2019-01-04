package main

import (
	"sync"
	"testing"
	"time"
	"unsafe"
)

type IntKey struct {
	v int
}

func (key *IntKey) LessThan(other Key) bool {
	return key.v < other.(*IntKey).v
}

func (key *IntKey) Equal(other Key) bool {
	return key.v == other.(*IntKey).v
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

func TestInsertAndDelete(t *testing.T) {
	list := NewList()
	wg := &sync.WaitGroup{}
	n := 1024 * 64
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			key := i % 2
			list.Insert(&IntKey{key})
			time.Sleep(time.Millisecond)
			list.Delete(&IntKey{key})
			wg.Done()
		}(i)
	}
	wg.Wait()
	if !((*Node)(list.head).next == list.tail) {
		t.Errorf("head.next must be tail")
	}
}
