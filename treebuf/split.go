package treebuf

import (
	"fmt"
//	"github.com/anlhord/generic"
)

func parent(min, max uintptr, now *Node) *Node {
	next := &now

	for {
		if min < now.Key && max < now.Key {
			next = &(now.l)
		} else if min > now.Key && max > now.Key {
			next = &(now.r)
		} else {
			if now.r == nil && now.l == nil {
print("DEAD\n")
				*next = nil
				return nil
			}

			return now
		}

		if *next == nil {
			return nil
		}

		node := **next
		node.p = now
		(*next) = &node
	 	now = &node

		fmt.Printf("nuke %d ~ %d nowkey %d\n", min, max, now.Key)
	}
}

func nukel(key uintptr, now *Node) *Node {
	if now == nil {
		return nil
	}

	other := &now
	next := &now

	for {
		if key == now.Key {
			*next = nil
			return now
		} else if (key > now.Key) {
			next = &(now.r)
			other = &(now.l)
		} else if (key < now.Key) {
			next = &(now.l)
			other = &(now.r)
		}

		if *next == nil {
			return nil
		}

		(*other) = nil
		node := **next
		node.p = now
		(*next) = &node
	 	now = &node

		fmt.Printf("nukeLL %d ~ nowkey %d\n", key, now.Key)
	}
}

// cele je to o tom r.trunk.l pointri
// treba ho na zaciatku skopirovat a potom atomicky vymenit
// avsak struktura v tom momente je nekonzistentna pretoze akonahle niekto
// prejde cez parent pointer tak moze vstupit do starej oblasti
// toto vyuziju mrtvi atteri ze presne po toto prijdu
// obnovenie konzistencie vyzaduje log(n) casu a realizuje sa druhym prechodom
// potom sa fixnu atteri a hotovo
func try_nuke(l **Node, min, max uintptr) bool {
// get the trunk.l in tmp
	tmp := Load(l)
	if tmp == nil {
		panic("????")
	}

	n := *tmp
	n.p = tmp


	p := parent(min, max, &n)
	if p == nil {
		print("\n\n\nDNO\n\n\n")
		return true
	}
	// p node must be deleted but his children preserved
	// need merge trees operator
	// najdem najpravejsie decko a dam ho tu
	//	*next = nil





	if p.l != nil {
	print("\n\n\n ma lave \n\n\n")
		foo := *(p.l)
		p.l = &foo
		nukel(min, &foo)
	}
	if p.r != nil {
	print("\n\n\n ma prave \n\n\n")
		foo := *(p.r)
		p.r = &foo
		nukel(min, &foo)
	}







/*
	else {

	}
*/
//	fmt.Printf("tmp %p\n", tmp)
	fmt.Printf("nuked %v\n", p)

	// now we need to 

// store n in trunk.l
	return CmpPutPtr(l, tmp, &n)
}

func (r *Root) Nuke(min, max uintptr) {
	try_nuke(&(r.trunk.l), min, max)
/*
	p := parent(min, max, )
	if p == nil {
		return
	}
	nuke_serial(min, max, p.l, p.r)
	_ = p
*/
}

// This is slow from non-root node. Use the above if possible
func (a *Atter) Nuke(min, max uintptr) {
	try_nuke(&(up(a.p).l), min, max)
/*
	p := parent(min, max, )
	if p == nil {
		return
	}
	nuke_serial(min, max, p.l, p.r)
	_ = p
*/
}
