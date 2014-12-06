package treebuf

import (
	"example.com/repo.git/colmgr"
//	"fmt"
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
	Nexter
}

type Nexter struct {
	p *Node	// this is never empty
}
func (n *Nexter) End() bool {
	return n.p == nil
}
func (n *Nexter) Next() {
	if n.p.r != nil {
		n.p = n.p.r
		for n.p.l != nil {
			n.p = n.p.l
		}
		return
	}
	// v pravo nic neni ideme hore

}
func (a *Atter) End() bool {
	// we are only end, when the tree is empty, or no smaller or equal key
	if a.Nexter.p.Key > a.key {
		return true
	}

	return a.Nexter.p.Trunk()
}

func (a *Atter) Next() colmgr.Nexter {	// we are only end, when the tree is empty
	n := &Nexter{p: a.Nexter.p}
	n.Next()
	return n
}

func (r *Root) At(key uintptr) colmgr.Atter {
	now := &r.trunk

	return &Atter{key:key, Nexter:Nexter{now}}
}
