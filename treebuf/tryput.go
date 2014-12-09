package treebuf

import (
	"github.com/anlhord/generic"
)


func (a *Atter) MkNode(key uintptr, val generic.Value) {
	if a.p.p == nil {
		if a.p.r == nil {
			a.p.r = node(key, val, a.p)
			return
		}
		mk(key, val, a.p.r)
		return
	}
	mk(key, val, a.p)
}
