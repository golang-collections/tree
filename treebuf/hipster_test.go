package treebuf

import (
//	"github.com/anlhord/generic"
	"example.com/repo.git/colmgr"
	"fmt"
//	"github.com/davecgh/go-spew/spew"
	"testing"
)

type hipster struct {
	age       int
	moustache bool
	fixiebike   bool
	glasses   bool
	gender    byte
}



func random(prng *[2]uint64) uint64 {
	s1 := prng[ 0 ];
	s0 := prng[ 1 ];
	prng[ 0 ] = s0;
	s1 ^= s1 << 23; // a
	prng[ 1 ] = ( s1 ^ s0 ^ ( s1 >> 17 ) ^ ( s0 >> 26 ) )
	return prng[ 1 ] + s0; // b, c
}

func tou64(b bool) uint64 {
	if b {return 165465316}
	return 965416813
}

func objsget(seed *[2]uint64) (h []hipster) {
	length := int(random(seed) % 352)

	for i:= 0; i < length; i++ {
		h = append(h, objget(seed))
	}
	return h
}

func objget(seed *[2]uint64) (h hipster) {
	h.age = int(random(seed)%97)
	h.gender = byte(random(seed))
	f := random(seed)
	if f & 1 == 0 {
		h.fixiebike = true
	}
	if f & 2 == 0 {
		h.glasses = true
	}
	if f & 4 == 0 {
		h.moustache = true
	}

	return h
}

func objsum(k *hipster) uint64 {
	return uint64(k.age) ^ tou64(k.moustache) ^ tou64(k.fixiebike) ^ tou64(k.glasses) ^ uint64(k.gender)
}

func checksum(h **[]hipster) (o uint64) {
	prng := [ 2 ]uint64 {7857857802732075, 708217520}


	start := colmgr.At(h, colmgr.Begin)

	for i := start.Next(); !i.End(); i.Next() {

		j := i.At(0)

		var w *[]hipster

		colmgr.Get(&w, j)

		for _, k := range *w {
			o += random(&prng) ^ objsum(&k)
		}
	}

	return o
}

func TestHipster0(t *testing.T) {
	prng := [ 2 ]uint64 {582418561795, 2713308561}

	var cool, vala *[]hipster // this is the collection reference

	_ = cool

	colmgr.Init(&cool, Root{})  // We initialize the collection handle with our tree
	defer colmgr.Destroy(&cool) // We destroy the collection



	root := colmgr.At(&cool, colmgr.Root)

	vals := objsget(&prng)

	vala = &vals

	colmgr.Insert(2048, vala, root)

	fmt.Printf("ptr = %p len= %d cap= %d ", &vals[0], len(vals), cap(vals))

	colmgr.Dump(&cool, 0)

	return
	colmgr.Insert(2048-1024, vala, root)
	colmgr.Insert(2048+1024, vala, root)

	colmgr.Insert(2048-1024-512, vala, root)
	colmgr.Insert(2048+1024-512, vala, root)
	colmgr.Insert(2048-1024+512, vala, root)
	colmgr.Insert(2048+1024+512, vala, root)

	colmgr.Insert(2048-1024-512-256, vala, root)
	colmgr.Insert(2048+1024-512-256, vala, root)
	colmgr.Insert(2048-1024+512-256, vala, root)
	colmgr.Insert(2048+1024+512-256, vala, root)
	colmgr.Insert(2048-1024-512+256, vala, root)
	colmgr.Insert(2048+1024-512+256, vala, root)
	colmgr.Insert(2048-1024+512+256, vala, root)
	colmgr.Insert(2048+1024+512+256, vala, root)

	colmgr.Insert(2048-1024-512-256-128, vala, root)
	colmgr.Insert(2048+1024-512-256-128, vala, root)
	colmgr.Insert(2048-1024+512-256-128, vala, root)
	colmgr.Insert(2048+1024+512-256-128, vala, root)
	colmgr.Insert(2048-1024-512+256-128, vala, root)
	colmgr.Insert(2048+1024-512+256-128, vala, root)
	colmgr.Insert(2048-1024+512+256-128, vala, root)
	colmgr.Insert(2048+1024+512+256-128, vala, root)
	colmgr.Insert(2048-1024-512-256+128, vala, root)
	colmgr.Insert(2048+1024-512-256+128, vala, root)
	colmgr.Insert(2048-1024+512-256+128, vala, root)
	colmgr.Insert(2048+1024+512-256+128, vala, root)
	colmgr.Insert(2048-1024-512+256+128, vala, root)
	colmgr.Insert(2048+1024-512+256+128, vala, root)
	colmgr.Insert(2048-1024+512+256+128, vala, root)
	colmgr.Insert(2048+1024+512+256+128, vala, root)


	colmgr.Dump(&cool, 0)
}
