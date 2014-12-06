package treebuf

import (
	"fmt"
)

// SCAFFOLDING OPERATORS ///////////////////////////////////////////////////////
func node(key uintptr, val []byte, p *Node) *Node {
	return &Node{r: nil, l: nil, p:p, Key:key, Val:val}
}

func mk(key uintptr, val []byte, now *Node) {
	next := now

	for {
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

func (r *Root) MkNode(key uintptr, val []byte) {

	fmt.Printf("MkNode to %d.\n", key)

	if r.trunk.r == nil {
		newval := node(key, val, nil)
		fmt.Printf("newwal to %d.\n", newval)
		if CmpSwapPtr(&(r.trunk.r), newval) {
			return
		}
	}



	mk(key, val, r.trunk.r)
}
