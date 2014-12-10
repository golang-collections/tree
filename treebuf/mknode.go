package treebuf

import (
	"fmt"
	"github.com/anlhord/generic"
)

func node(key uintptr, val generic.Value, p *Node) *Node {
	return &Node{r: nil, l: nil, p: p, Key: key, Val: val}
}

func mk(key uintptr, val generic.Value, now *Node) {
	next := now

	for {
//		fmt.Printf("key %d nowkey %d\n", key, now.Key)

		if key < now.Key {
			next = now.l
			if next == nil {
				if CmpSwapPtr(&now.l, node(key, val, now)) {
					return
				} else {
					continue
				}
			}
		} else if key > now.Key {
			next = now.r
			if next == nil {
				if CmpSwapPtr(&now.r, node(key, val, now)) {
					return
				} else {
					continue
				}
			}
		} else {
			return
		}

		now = next
	}
}

func mkup(key uintptr, n *Node) *Node {
	for n.p != nil && ((n.Key < n.p.Key && key >= n.p.Key) || (n.Key >= n.p.Key && key < n.p.Key)) {
//		fmt.Printf("MKUP from %p to %p.\n", n, n.p)
		n = n.p
	}
	return n
}

func (r *Root) MkNode(key uintptr, val generic.Value) {

//	fmt.Printf("MkNode to %d.\n", key)

	mk(key, val, &r.trunk)
}


// If node truly is in At node's subtree use this
// This operator is UNSAFE and destroys the tree if the precondition is not met.
func (a *Atter) MkNode(key uintptr, val generic.Value) {
	fmt.Printf("MkNode to %d from %d.\n", key, a.p.Key)

	mk(key, val, mkup(key, a.p))
}
