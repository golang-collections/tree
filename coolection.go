package main

import (
        "fmt"
//        "unsafe"
	"reflect"
 //       "math/rand"
)

type Config struct {	// this will be inside the iterator& Root
	lsum bool	// the structure operates in a left subtree sum mode
}
type Collector interface {
	Spawner
}

type Spawner interface {
	Spawn() Cursor
}

type Initer interface {
	Init() Collector
}

// coolection manager

var EEnd = fmt.Errorf("End of collection")
var coolections map[uintptr]Collector

func init() {
	coolections = make(map[uintptr]Collector)
}

func Init(handle interface{}, initer Initer) {
	p := uintptr(reflect.ValueOf(handle).Pointer())

	fmt.Printf("Calling initer to %d.\n", p)


}

type Cursor interface {
	Next() (Cursor)
	Prev() (Cursor)

}

func Destroy(handle interface{}) {
	p := uintptr(reflect.ValueOf(handle).Pointer())
	coolections[p] = nil

	fmt.Printf("Destroyed %d.\n", p)
	// FIXME: refcounting?
}

func Spawn(handle interface{}) Cursor {
	p := uintptr(reflect.ValueOf(handle).Pointer())

	obj := coolections[p]

	fmt.Printf("Spawning %d on %v.\n", p, obj)

	return obj.Spawn()
}
