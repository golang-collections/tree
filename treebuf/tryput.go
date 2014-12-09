package treebuf

import (
	"github.com/anlhord/generic"
)

func (a *Atter) MkNode(key uintptr, val generic.Value) {
	mk(key, val, a.p)
}
