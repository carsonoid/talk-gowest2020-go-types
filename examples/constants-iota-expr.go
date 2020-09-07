package main

import "fmt"

// BEGIN OMIT
const (
	x = iota * 2     // index 0 means iota == 0
	y = iota + iota  // index 1 means iota == 1
	_                // index 2 means iota == 2
	z = iota * 1.002 // index 3 means iota == 3
)

func main() {
	fmt.Println(x, y, z)
}

// END OMIT
