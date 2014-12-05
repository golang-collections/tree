package main


import (
 //       "fmt"
 //       "math/rand"
)

// this is a stupid tree

type Lambda struct { // iterator will carry this
	pos	*Node
}

func (l Lambda) Next() Cursor {
	return l
}

func (l Lambda) Prev() Cursor {
	return l
}

type Ender struct {
	Lambda
}

func (Ender) Next() Cursor {
	return Ender{}
}
func (e Ender) Prev() Cursor {
	if e.pos == nil {
		return Ender{}
	} else {
		return Lambda{e.pos}
	}
}

type Root struct {
	r *Node
	size uintptr	// this is the total size of the arraytree
}

type Node struct {
	l, r, u *Node
	Key	uintptr		// this is the key
	Val	[]byte		// this is the val
}

func (Root) Init() Collector {


	return &Root{}
}

func (r Root) Spawn() Cursor {

	print("spawning cursor on root\n")

	if r.r == nil {
		return Ender{}
	}

	var l Lambda
	l.pos = r.r
	return l
}
