package treemap

type Node struct {
	r, l, p *Node
	key     uintptr
	val     []byte
}
type Root struct {
	root *Node
}

func node(key uintptr, p *Node) *Node {
	return &Node{r: nil, l: nil, p:p, key:key, val:nil}
}

// SCAFFOLDING OPERATORS ///////////////////////////////////////////////////////

func Mk(key uintptr, now *Node) {
	next := now

	for {
		if key < now.key {
			next = now.l
			if next == nil {
				now.l = node(key, now)
				return
			}
		} else if key > now.key {
			next = now.r
			if next == nil {
				now.r = node(key, now)
				return
			}
		} else {
			return
		}

		now = next
	}
}

func MkNode(key uintptr, root *Root) {
	now := root.root

	if now == nil {
		root.root = node(key, nil)
	}
	Mk(key, now)
}
