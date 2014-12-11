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


	root := colmgr.At(&cool, colmgr.Root)

	// We put values to a collection

	root.MkNode(512, vala)

	root.MkNode(256, vala)
	root.MkNode(768, vala)

	root.MkNode(128, vala)
	root.MkNode(384, vala)
	root.MkNode(640, vala)
	root.MkNode(896, vala)


	// We create iterators at various spots in the collection

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

		j := i.At(0)

		window := j.Map()

		fmt.Printf("%s", window)

		j.Upd(valb)

		i = j.Next()

	}

	for i := start.Next(); !i.End(); i.Next() {
		j := i.At(0)

		window := j.Map()

		fmt.Printf("%s", window)

		i = j.Next()
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

func TestTryFreezeNexter0(t *testing.T) {
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
		j := i.At(0)	// freeze nexter on the current slice 0th object
		_ = j

		window := j.Map()
		fmt.Printf("%s", window)


		ni := j.Next()	// unfreeze a nexter

		foo := fmt.Sprint("%v", i)
		bar := fmt.Sprint("%v", ni)


		if foo != bar {
			t.Fatal("next is not same:", i, ni)
		}
	}
}

func TestTryBadMkNode0(t *testing.T) {
	// This test demonstrates how a correct tree is produced
	// by a correct use of Atter.MkNode()

	var cool generic.Collection // this is the collection reference
	colmgr.Init(&cool, Root{})  // We initialize the collection handle with our tree
	defer colmgr.Destroy(&cool) // We destroy the collection

	root := colmgr.At(&cool, colmgr.Root) // get root node Atter
	_ = root

	vala := []byte("bad")
	valb := []byte("good")

	_ = vala
	_ = valb

	root.MkNode(512, valb)

	root.MkNode(256, valb)
	root.MkNode(768, valb)



	leftsubtree := root.At(256)
	rightsubtree := root.At(768)

//	this is ok
//	spew.Dump("lstree:", leftsubtree)

	rightsubtree.MkNode(128, vala)
	rightsubtree.MkNode(384, vala)
	leftsubtree.MkNode(640, vala)
	leftsubtree.MkNode(896, vala)

	print("---------------\n")

	colmgr.Dump(&cool, 0)

	print("---------------\n")
}

func TestTryBadMkNode1(t *testing.T) {
	// This test demonstrates how a broken tree is produced
	// by an incorrect use of Atter.MkNode()

	var cool generic.Collection // this is the collection reference
	colmgr.Init(&cool, Root{})  // We initialize the collection handle with our tree
	defer colmgr.Destroy(&cool) // We destroy the collection

	root := colmgr.At(&cool, colmgr.Root) // get root node Atter
	_ = root

	vala := []byte("bad")
	valb := []byte("good")

	_ = vala
	_ = valb

	root.MkNode(512, valb)
	root.MkNode(599, valb)
	root.MkNode(513, valb)
	root.MkNode(598, valb)

	// todo put 514

	root.MkNode(2, valb)
	root.MkNode(511, valb)
	root.MkNode(3, valb)
	root.MkNode(510, valb)

	nower := root.At(510)

//	print("---------------\n")

//	colmgr.Dump(&cool, 0)

	nower.MkNode(514, vala)

//	print("---------------\n")

//	colmgr.Dump(&cool, 0)

//	print("---------------\n")
}

func TestTryAppend0(t *testing.T) {
	// This test demonstrates how a broken tree is produced
	// by an incorrect use of Atter.MkNode()

	var cool generic.Collection // this is the collection reference
	colmgr.Init(&cool, Root{})  // We initialize the collection handle with our tree
	defer colmgr.Destroy(&cool) // We destroy the collection

	root := colmgr.At(&cool, colmgr.Root) // get root node Atter
	_ = root

	vala := []byte("X")
	_ = vala

	end := root.At(colmgr.End)
	_ = end


	for i := 0 ; i < 10 ; i++ {
		end.Append(vala)
		end.Fix()
	}

	print("---------------\n")

	colmgr.Dump(&cool, 0)

	for i := 0 ; i < 10 ; i++ {
		root.Append(vala)
	}

	print("---------------\n")

	colmgr.Dump(&cool, 0)

}
