package treebuf

import (
	"github.com/anlhord/generic"
	"github.com/anlhord/generic/length"
)

// Push is only possible to the last node.
func (a *Atter) Push(val *generic.Value) {
	now := a.p
	if a.p.p == nil {
		if a.p.l == nil {



			a.MkNode(0, *val)
			return
		}
		now = now.l
	}
	for now.r != nil {
		now = now.r
	}
	l := length.Len(now.Val)

	mk(now.Key+uintptr(l), *val, now)
}
