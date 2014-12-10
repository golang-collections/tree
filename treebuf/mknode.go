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
		fmt.Printf("key %d nowkey %d\n", key, now.Key)

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

// SCAFFOLDING OPERATOR
func (r *Root) MkNode(key uintptr, val generic.Value) {

	fmt.Printf("MkNode to %d.\n", key)

	if r.trunk.r == nil {
		newval := node(key, val, &r.trunk)
		fmt.Printf("newwal to %d.\n", newval)
		if CmpSwapPtr(&(r.trunk.r), newval) {
			return
		}
	}

	mk(key, val, r.trunk.r)
}

func (a *Atter) MkNode(key uintptr, val generic.Value) {
	fmt.Printf("MkNode to %d from %d.\n", key, a.p.Key)

	if a.p.p == nil {
		if a.p.r == nil {
			a.p.r = node(key, val, a.p)
			return
		}
		mk(key, val, a.p.r)
		return
	}
	mk(key, val, a.p)
}
