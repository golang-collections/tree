package treebuf

import (
	"sync/atomic"
	"unsafe"
)

/*
        atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&head.nxt)), unsafe.Pointer(head.nxt), unsafe.Pointer(n))
*/


func CmpSwapPtr(addr **Node, val *Node) (swapped bool) {
	if !thread_safe_mknod {
		(*addr) = val
		return true
	}

	p := (*unsafe.Pointer)(unsafe.Pointer(addr))
	return atomic.CompareAndSwapPointer(p, nil, unsafe.Pointer(val))
}

func CmpPutPtr(addr **Node, old *Node, val *Node) (swapped bool) {
	if !thread_safe_nuke {
		(*addr) = val
		return true
	}

	p := (*unsafe.Pointer)(unsafe.Pointer(addr))
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(val))
}

func Load(addr **Node) *Node {
	if !thread_safe_nuke {
		return *addr
	}
	p := (*unsafe.Pointer)(unsafe.Pointer(addr))
	return (*Node)(atomic.LoadPointer(p))
}
