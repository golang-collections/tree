package treebuf

import (
	"example.com/repo.git/colmgr"
	"fmt"
)

type Root struct {
	trunk Node
//	RWMutex		// not needed currently
}

type Node struct {
	l, r, p *Node
	Key	uintptr		// this is the key
	Val	[]byte		// this is the val
}
func (Root) Root() colmgr.Collector {
	return &Root{}
}
func (n *Node) Trunk() bool {
	return n.p == nil
}

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
	if debug_destructor && r.trunk.l == &r.trunk {
		panic("Dead tree")
	}

	now := &r.trunk

	return &Atter{key:key, p:now}
}

// XXX /////////////////////////////////////////////////////////////////////////

func (r *Node) Dump(f byte, d uint) {
	if r.l != nil {
		r.l.Dump(f, d+1)
	}
	for i := uint(0) ; i < d; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("[%v]\n", r)
	if r.r != nil {
		r.r.Dump(f, d+1)
	}
}

func (r *Root) Dump(f byte) {
	if debug_destructor && r.trunk.l == &r.trunk {
		panic("Dead tree")
	}
	fmt.Printf("Dumping the tree %p with format %d \n", r, f)
	if r.trunk.r != nil {
		r.trunk.r.Dump(f, 0)
	}
}

// XXX /////////////////////////////////////////////////////////////////////////

func (r *Root) Destroy() {
	if debug_destructor {
		r.trunk.l = &r.trunk
	}
	r.trunk.r = nil
	// FIXME: more cleanup?
}
