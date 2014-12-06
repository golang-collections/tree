package treebuf

import (
	"sync/atomic"
	"unsafe"
)

func CmpSwapPtr(addr **Node, val *Node) (swapped bool) {
	if !thread_safe_mknod {
		(*addr) = val
		return true
	}
	p := unsafe.Pointer(*addr)
	return atomic.CompareAndSwapPointer(&p, nil, unsafe.Pointer(val))
}
