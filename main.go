package main

import (
	"fmt"
	"reflect"
)

type Ring struct {
	size int
}

func main() {

	var servers = make(map[string]int, 0)

	fmt.Println(reflect.TypeOf(servers))

	ring := NewRing(10)
	fmt.Println(ring)

	// add servers to the ring
	servers["server1"] = 1
	servers["server2"] = 2

	fmt.Println(servers)

}

func NewRing(size int) *Ring {
	return &Ring{
		size: size,
	}
}
