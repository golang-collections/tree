package treebuf

import "example.com/repo.git/colmgr"

// ITERATORS STUFF /////////////////////////////////////////////////////////////
type Atter struct {
	key uintptr	// we try to stay at or after this key
	p *Node		// never nil
}

type Nexter struct {
	p, q *Node	// p is never nil
}
func (n *Nexter) End() bool {
	return n.q == nil
}
func (n *Nexter) gap() uintptr {
	return n.q.Key - n.p.Key
}
func (n *Nexter) Next() {
	n.p = n.q
	if n.q.r != nil {
		n.q = n.q.r
		for n.q.l != nil {
			n.q = n.q.l
		}
		return
	}
	// v pravo nic neni ideme hore

}
func (a *Atter) End() bool {
	// we are only end, when the tree is empty, or no smaller or equal key
	if a.p.Key > a.key {
		return true
	}

	return a.p.Trunk()
}

func (a *Atter) Next() colmgr.Nexter {	// we are only end, when the tree is empty
	n := &Nexter{p: a.p, q: a.p}
	n.Next()
	return n
}

func (r *Root) At(key uintptr) colmgr.Atter {
	now := &r.trunk

	return &Atter{key:key, p:now}
}
