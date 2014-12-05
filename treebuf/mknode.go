package treebuf

// SCAFFOLDING OPERATORS ///////////////////////////////////////////////////////
func node(key uintptr, p *Node) *Node {
	return &Node{r: nil, l: nil, p:p, Key:key, Val:nil}
}

func mk(key uintptr, now *Node) {
	next := now

	for {
		if key < now.Key {
			next = now.l
			if next == nil {
				if CmpSwapPtr(&now.l, node(key, now)) {
					return
				} else {
					continue
				}
			}
		} else if key > now.Key {
			next = now.r
			if next == nil {
				if CmpSwapPtr(&now.r, node(key, now)) {
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

func (root *Root) MkNode(key uintptr) {
	if root.r == nil {
		root.r = node(key, nil)
		return
	}

//	fmt.Printf("MkNode to %d.\n", key)

	mk(key, root.r)
}
