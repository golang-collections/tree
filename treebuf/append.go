package treebuf

import (
	"github.com/anlhord/generic"
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
	mk(now.Key+uintptr(len(now.Val)), *val, now)
}
