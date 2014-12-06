package treebuf

import (
	"example.com/repo.git/colmgr"
	"fmt"
)

// ITERATORS STUFF /////////////////////////////////////////////////////////////
type Atter struct {
	key uintptr // we try to stay at or after this key
	p   *Node   // never nil
}

type Nexter struct {
	p, q *Node // p is never nil
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
	fmt.Println("v pravo nic neni ideme hore\n")
	// v pravo nic neni ideme hore

}
func (a *Atter) End() bool {
	// we are only end, when the tree is empty, or no smaller or equal key
	if a.p.Key > a.key {
		return true
	}

	return a.p.Trunk()
}

func (a *Atter) Next() colmgr.Nexter { // we are only end, when the tree is empty
	fmt.Println("Idem dalej s attera\n")

	n := &Nexter{p: a.p, q: a.p}
	fmt.Printf("som %p %p \n", n.p, n.q)
	n.Next()
	fmt.Printf("som %p %p \n", n.p, n.q)
	return n
}

func (r *Root) At(key uintptr) colmgr.Atter {
	if r.trunk.r == nil {
		return &Atter{key: key, p: &r.trunk}
	}
	now := r.trunk.r
	ok := now

//	for now.l != nil && now.r != nil {
	for now.r != nil && now.Key < key {
	fmt.Printf(".. key=%d %p\n", key, now)
		ok = now
		now = now.r
	}

	for now.l != nil && now.Key > key {
	fmt.Printf(",, key=%d %p -> %p\n", key, now, now.l)
		now = now.l
	}

	if now.Key > key && ok.Key < now.Key {
	fmt.Printf("! key=%d %p -> %p\n", key, now, ok)
		now = ok

	} 

//	}
	fmt.Printf("Atol som sa na key=%d %p\n", key, now)
	return &Atter{key: key, p: now}
}
