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

	mk(key, val, &r.trunk)
}

func (a *Atter) MkNode(key uintptr, val generic.Value) {
	fmt.Printf("MkNode to %d from %d.\n", key, a.p.Key)

	mk(key, val, a.p)
}
