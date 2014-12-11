package treebuf

import (
	"github.com/anlhord/generic"
)

// Append is only possible to the last node.
func (a *Atter) Append(val generic.Value) {
	if a.p.p == nil {
		a.MkNode(0, val)
		return
	}
	if a.p.r == nil {
		a.MkNode(a.p.Key + uintptr(len(a.p.Val)), val)
		return
	}
	panic("append only to the end")
}
