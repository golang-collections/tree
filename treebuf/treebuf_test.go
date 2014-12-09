package treebuf

import (
	"example.com/repo.git/colmgr"
	"fmt"
	"github.com/anlhord/generic"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestSeekStart0(t *testing.T) {
	spew.Dump("")
}

func TestMy0(t *testing.T) {
	var cool generic.Collection // this is the collection reference
	colmgr.Init(&cool, Root{})  // We initialize the collection handle with our tree
	defer colmgr.Destroy(&cool) // We destroy the collection

	vala := []byte("hello")
	valb := []byte("world")

	// We put values to a collection

	colmgr.MkNode(&cool, 512, vala)

	colmgr.MkNode(&cool, 256, vala)
	colmgr.MkNode(&cool, 768, vala)

	colmgr.MkNode(&cool, 128, vala)
	colmgr.MkNode(&cool, 384, vala)
	colmgr.MkNode(&cool, 640, vala)
	colmgr.MkNode(&cool, 896, vala)

	// We create iterators at various spots in the collection

	root := colmgr.At(&cool, colmgr.Root)
	_ = root.At(767)
	at := root.At(768)
	near := root.At(769)
	start := root.At(colmgr.Begin)
	end := root.At(colmgr.End)
	_ = at
	_ = near
	_ = start
	_ = end
	_ = root

	//	spew.Dump(root)

	colmgr.Dump(&cool, 0)

	for i := start.Next(); !i.End(); i.Next() {

		window := i.Map()

		fmt.Printf("%s", window)

		i.Upd(valb)

	}

	for i := start.Next(); !i.End(); i.Next() {
		window := i.Map()

		fmt.Printf("%s", window)
	}

	// now we loop around the collection and put an actual values
	// from start to end
	/*

		// print the len
		print(len(cool))

		// do a subslice
		foo := slice(cool, 50, 100)

		for i, j := range cool {
			print(i, j)
			//do a slow operation
		}

		cool = append(cool, []byte("foo"))
	*/

	//	next_pointer :=

}

func TestTryPut0(t *testing.T) {
	var cool generic.Collection // this is the collection reference
	colmgr.Init(&cool, Root{})  // We initialize the collection handle with our tree
	defer colmgr.Destroy(&cool) // We destroy the collection

	root := colmgr.At(&cool, colmgr.Root) // get root node Atter
	_ = root

	vala := []byte("hello")
	valb := []byte("world")

	_ = vala
	_ = valb

	root.MkNode(512, vala)

	root.MkNode(256, vala)
	root.MkNode(768, vala)

	leftsubtree := root.At(256)
	rightsubtree := root.At(768)

	leftsubtree.MkNode(128, vala)
	leftsubtree.MkNode(384, vala)
	rightsubtree.MkNode(640, vala)
	rightsubtree.MkNode(896, vala)

	begin := root.At(colmgr.Begin)

	colmgr.Dump(&cool, 0)

	for i := begin.Next(); !i.End(); i.Next() {
		window := i.Map()
		fmt.Printf("%s", window)


		j := i.At(0)	// freeze nexter on the current slice 0th object
		_ = j
		ni := j.Next()	// unfreeze a nexter

		foo := fmt.Sprint("%v", i)
		bar := fmt.Sprint("%v", ni)


		if foo != bar {
			t.Fatal("next is not same:", i, ni)
		}
	}
}

/*
func TestAppend1(t *testing.T) {

	var h [1][]byte

	Append(&h, byte(254))
	if Checksum(h) != 36342608889142753 {
		t.Error("Simple byte append", Checksum(h))
	}
}

func TestAppend2(t *testing.T) {
	var h [1][]int
	mmm := [3]int{1, 2, 3}

	if ChecksumInt(h) != 2166136261 {
		t.Error("Initial append")
	}
	Appends(&h, [1]int{0})
	if ChecksumInt(h) != 36342608889142559 {
		t.Error("0 int append")
	}
	Appends(&h, mmm)
	if ChecksumInt(h) != 10900344298869587409 {
		t.Error("123 int append")
	}
	Append(&h, 4, 5, 6)
	if ChecksumInt(h) != 877089633685397766 {
		t.Error("456 int append")
	}
}
*/
