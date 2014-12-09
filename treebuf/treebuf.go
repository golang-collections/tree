package treebuf

import (
	"example.com/repo.git/colmgr"
	"fmt"
	"github.com/anlhord/generic"
)

type Root struct {
	trunk Node
}

type Node struct {
	l, r, p *Node
	Key     uintptr       // this is the key
	Val     generic.Value // this is the val
}

func (Root) Root() colmgr.Collector {
	return &Root{trunk: Node{Key: colmgr.Root, Val: nil}}
}
func (n *Node) Trunk() bool {
	return n.p == nil
}

// XXX /////////////////////////////////////////////////////////////////////////

func (r *Node) Dump(f byte, d uint) {
	if r.l != nil {
		r.l.Dump(f, d+1)
	}
	for i := uint(0); i < d; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("[%p | %v]\n", r, r)
	if r.r != nil {
		r.r.Dump(f, d+1)
	}
}

func (r *Root) Dump(f byte) {
	fmt.Printf("Dumping the tree %p with format %d \n", r, f)
	if r.trunk.r != nil {
		r.trunk.r.Dump(f, 0)
	}
}

// XXX /////////////////////////////////////////////////////////////////////////

func (r *Root) Destroy() {
	if debug_destructor {
		if r.trunk.p != nil || r.trunk.l != nil || r.trunk.Key != colmgr.Root {
			panic("Dubious trunk")
		}
		r.trunk.l = &r.trunk
	}
	r.trunk.r = nil
	// FIXME: more cleanup?
}
