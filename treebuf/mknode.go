package treebuf

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

func (root *Root) MkNode(key uintptr, val []byte) {
	if root.trunk.r == nil {
		if CmpSwapPtr(&root.trunk.r, node(key, val, nil)) {
			return
		}
	}

//	fmt.Printf("MkNode to %d.\n", key)

	mk(key, val, root.trunk.r)
}
