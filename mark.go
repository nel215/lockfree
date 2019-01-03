package main

import (
	"unsafe"
)

func mark(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) | uintptr(1))
}

func unmark(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) & (^uintptr(1)))
}

func marked(p unsafe.Pointer) bool {
	return (uintptr(p) & uintptr(1)) == uintptr(1)
}
