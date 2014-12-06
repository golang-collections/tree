package treebuf

import (
	"example.com/repo.git/colmgr"
//	"fmt"
)

type Root struct {
	r *Node
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
type Ender struct {}

func (r *Root) MkSeek() colmgr.Seeker {
	if r.r == nil {
		return Ender{}
	}
	return 
}
