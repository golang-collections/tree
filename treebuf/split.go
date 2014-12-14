package treebuf

import (
	"fmt"
//	"github.com/anlhord/generic"
)

func parent(min, max uintptr, now *Node) *Node {
	next := now

	for {
		if min < now.Key && max < now.Key {
			next = now.l
		} else if min > now.Key && max > now.Key {
			next = now.r
		} else {
		fmt.Printf("nuke %d ~ %d nowkey %d\n", min, max, now.Key)

			return now
		}
		if next == nil {
			return nil
		}
		now = next
	}
}

func nuke_serial(min, max uintptr, l, r *Node) {
	for l.l != nil && r.r != nil {
		if min < l.Key && l.l != nil {

			fmt.Printf("nukel %d nowkey %d\n", min, l.Key)

			l = l.l

		}

		if max > r.Key && r.r != nil {

			fmt.Printf("nuker %d nowkey %d\n", min, r.Key)

			r = r.r

		}
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



	fmt.Printf("tmp %p\n", tmp)
	fmt.Printf("nuker %v\n", n)


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
