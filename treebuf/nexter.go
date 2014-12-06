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

	fmt.Printf("som %p %p \n", n.p, n.q)

	if n.q.l == nil && n.q.r == nil {
		t := n.q
		n.q = n.p
		n.p = t

		return
	}




//	for n.q.r == n.p || n.q.l == n.p {
	// vyliezam z prava
	if n.q.r == n.p {
		n.p = n.q
		n.q = n.q.p
//		n.Next()
		return
	}
	// vyliezam z lava
	if n.q.l == n.p {
		n.p = n.q
		n.q = n.q.r
		n.Next()
		return
	}
//	}



	// lezem do lava
	for n.q.l != nil {
		n.p = n.q
		n.q = n.q.l
	}
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

	q := a.p

	if q.l != nil {
		q = q.l
	} else if q.r != nil {
		q = q.r
	} else {
		q = q.p
	}

	n := &Nexter{p: a.p, q: q}
	return n
}

func (r *Root) At(key uintptr) colmgr.Atter {
	if r.trunk.r == nil {
		return &Atter{key: key, p: &r.trunk}
	}
	now := r.trunk.r
	ok := now

	for now.r != nil && now.Key < key {
//	fmt.Printf(".. key=%d %p\n", key, now)
		ok = now
		now = now.r
	}

	for now.l != nil && now.Key > key {
//	fmt.Printf(",, key=%d %p -> %p\n", key, now, now.l)
		now = now.l
	}

	if now.Key > key && ok.Key < now.Key {
//	fmt.Printf("! key=%d %p -> %p\n", key, now, ok)
		now = ok

	} 

//	fmt.Printf("Atol som sa na key=%d %p\n", key, now)
	return &Atter{key: key, p: now}
}
