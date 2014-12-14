package treebuf

import (
	"fmt"
	"github.com/anlhord/generic"
)

func node(key uintptr, val *generic.Value, p *Node) *Node {
	return &Node{r: nil, l: nil, p: p, Key: key, Val: *val}
}

func mk(key uintptr, val *generic.Value, now *Node) {
	next := &now

	for {
		if debug_mknode {
			fmt.Printf("key %d nowkey %d\n", key, now.Key)
		}

		if key < now.Key {
			next = &(now.l)
		} else if key > now.Key {
			next = &(now.r)
		} else {
			return
		}
		if *next == nil {
			if CmpSwapPtr(next, node(key, val, now)) {
				return
			} else {
				continue
			}
		}
		now = *next
	}
}

func up(n *Node) *Node {
	for n.p != nil {
		n = n.p
	}
	return n
}

func mkup(key uintptr, n *Node) *Node {
	for n.p != nil && ((n.Key < n.p.Key && key >= n.p.Key) || (n.Key >= n.p.Key && key < n.p.Key)) {
		if debug_mknode {
			fmt.Printf("MKUP from %p to %p.\n", n, n.p)
		}
		n = n.p
	}
	return n
}

func (r *Root) MkNode(key uintptr, val *generic.Value) {
	if debug_mknode {
		fmt.Printf("MkNode to %d.\n", key)
	}
	mk(key, val, &r.trunk)
}

// If node truly is in At node's subtree use this
// This operator is UNSAFE and destroys the tree if the precondition is not met.
func (a *Atter) MkNode(key uintptr, val *generic.Value) {
	if debug_mknode {
		fmt.Printf("MkNode to %d from %d.\n", key, a.p.Key)
	}
	mk(key, val, mkup(key, a.p))
}
