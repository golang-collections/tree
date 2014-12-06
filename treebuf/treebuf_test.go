package treebuf

import (
	"testing"
	"example.com/repo.git/colmgr"
)

func TestSeekStart0(t *testing.T) {

}

func TestMy0(t *testing.T) {
	var cool [0][]byte	// this is the collection reference
	colmgr.Init(&cool, Root{})	// We initialize the collection
	defer colmgr.Destroy(&cool)	// We destroy the collection

	vala := []byte("hello")

	colmgr.MkNode(&cool, 512, vala)

	colmgr.MkNode(&cool, 256, vala)
	colmgr.MkNode(&cool, 768, vala)

	colmgr.MkNode(&cool, 128, vala)
	colmgr.MkNode(&cool, 384, vala)
	colmgr.MkNode(&cool, 640, vala)
	colmgr.MkNode(&cool, 896, vala)

	at := colmgr.At(&cool, 512)
	near := colmgr.At(&cool, 513)
	_ = at
	_ = near



	// now we loop around the collection and put an actual values
	// from start to end
/*
	print(len(cool))

	for i, j := range cool {
		print(i, j)
		//do a slow operation
	}

	cool = append(cool, []byte("foo"))
*/

//	next_pointer := 
	
}

func TestAppend0(t *testing.T) {
/*
	var h [1][]byte

	Append(&h, byte(19), byte(5), byte(4), byte(3))
	if Checksum(h) != 4433864318700699030 {
		t.Error("Variadic append")
	}
*/
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
