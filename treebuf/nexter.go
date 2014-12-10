package treebuf

import (
	"example.com/repo.git/colmgr"
	"github.com/anlhord/generic"
//	"fmt"
)

// ITERATORS STUFF /////////////////////////////////////////////////////////////
type Atter struct {
	key uintptr // we try to stay at or after this key
	p   *Node   // never nil
}

type NexterAtter struct {
	Atter
	direction byte
}
func (a *NexterAtter) Map() generic.Value {
	return a.Atter.Map()
}

func (a *NexterAtter) Upd(b generic.Value) {
	a.Atter.Upd(b)
}
func (a *NexterAtter) Fix() {
	a.Atter.Fix()
}

func (a *NexterAtter) At(key uintptr) colmgr.Atter {
	return a.Atter.At(key)
}
func (a *NexterAtter) End() bool {
	return a.Atter.End()
}

func (a *NexterAtter) Next() colmgr.Nexter { // we are only end, when the tree is empty
	n := (&a.Atter).next()
	switch a.direction {
	case 1:
		n.q = n.p.l
	case 2:
		n.q = n.p.r
	case 3:
		n.q = n.p.p
	default:
		n.q = nil
	}

	return n
}
type Nexter struct {
	p, q *Node // p is never nil
}

// elem is the element you want to fix on in a slice
func (n *Nexter) At(elem uintptr) colmgr.Atter {
	var dir byte

	switch n.q {
	case n.p.l:
		dir = 1
	case n.p.r:
		dir = 2
	case n.p.p:
		dir = 3
	}

	return &NexterAtter{Atter:Atter{key: n.p.Key + elem, p: n.p}, direction: dir}
}
func (n *Nexter) End() bool {
	return n.q == nil
}
func (n *Nexter) Map() generic.Value {
	return n.p.Val
}

func (n *Nexter) Upd(b generic.Value) {
	if uint(len(b)) > uint(n.gap()) {
		print("\nExceeded gap len.\n")
	}
	n.p.Val = b
}
func (a *Atter) Map() generic.Value {
	return a.p.Val
}

func (a *Atter) Upd(b generic.Value) {
}
func (n *Nexter) gap() uintptr {
	return n.q.Key - n.p.Key
}
func (n *Nexter) Next() {
	//	fmt.Printf("som %p %p \n", n.p, n.q)

	// lezem do lava
	for n.p.l == n.q || n.p.r == n.q {
		if n.q.l != nil {
			n.p = n.q
			n.q = n.q.l
		} else if n.q.r != nil {
			n.p = n.q
			n.q = n.q.r
		} else {
			break
		}
	}

	if n.q.l == nil && n.q.r == nil {
		t := n.q
		n.q = n.p
		n.p = t

		return
	}

	//	for n.q.r == n.p || n.q.l == n.p {
	// vyliezam z prava
	for n.q.r == n.p && n.q.p != nil {
		n.p = n.q
		n.q = n.q.p
		if n.q.l == n.p {
			n.p = n.q
			n.q = n.q.r
			return
		}
	}
	if n.q.r == n.p && n.q.p == nil {
		n.p = n.q
		n.q = nil
		return
	}

	// vyliezam z lava
	if n.q.l == n.p {
		n.p = n.q
		n.q = n.q.r
		return
	}
	//	}

}
func (a *Atter) Fix() {
	for a.p.l == a.p.r {
		a.p = a.p.l
	}
}

// FIXME: At() from non-root node is slow
func (a *Atter) At(key uintptr) colmgr.Atter {
	now := at(key, up(a.p))

//	fmt.Printf("Atol som sa na key=%d %p\n", key, now)
	return &Atter{key: key, p: now}
}
func (a *Atter) End() bool {
	return a.p.Trunk()
}

func (a *Atter) next() *Nexter {
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

func (a *Atter) Next() colmgr.Nexter {
	return (a).next()
}

func (r *Root) At(key uintptr) colmgr.Atter {
	if r.trunk.r == nil || key == colmgr.Root {
		return &Atter{key: key, p: &r.trunk}
	}
	now := at(key, r.trunk.r)

//	fmt.Printf("Atol som sa na key=%d %p\n", key, now)
	return &Atter{key: key, p: now}
}
func up(now *Node) *Node {
	for now.p != nil {
		now = now.p
	}
	return now
}

func at(key uintptr, now *Node) *Node {
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
	return now
}
