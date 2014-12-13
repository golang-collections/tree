package treebuf

func (r *Root) Split(min, max uintptr) {

}

func parent(min, max uintptr, now *Node) *Node {
	next := now

	for {
		//		fmt.Printf("key %d nowkey %d\n", key, now.Key)

		if min < now.Key && max < now.Key {
			next = now.l
		} else if min > now.Key && max > now.Key {
			next = now.r
		} else {
			return now
		}
		if next == nil {
			print("wtf")
		}
		now = next
	}
}

func (a *Atter) Split(min, max uintptr) {

}
