package main

func main() {
	var a [0][]byte	// this is a collection handle. You pass it by reference to users. When not needed, you destroy it.
	Init(&a, Root{})

	var b Cursor
	b = Spawn(&a)	// spawn a cursor at the collection root
	_ = b

	// TODO: find element at position in arraytree


	// TODO:



/*
	Insert(0, &interator, []byte("HELLO"))
	Insert(100, &interator, []byte("go"))
	Insert(200, &interator, []byte("WoRlD"))
	Insert(300, &interator, []byte("ghqear"))
	Insert(400, &interator, []byte("ahdsdaga"))
	Insert(500, &interator, []byte("hathaer5"))
	Insert(600, &interator, []byte("ysrhsdfh"))
*/


	Destroy(&a)
}
