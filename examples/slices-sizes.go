package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []bool{}
	fmt.Println("header bytes: ", unsafe.Sizeof(s1))
	fmt.Println(
		"array bytes: ", cap(s1),
	)

	fmt.Println()

	s2 := make([]bool, 100)
	fmt.Println("header bytes: ", unsafe.Sizeof(s2))
	fmt.Println(
		"array bytes: ", cap(s2),
	)
}
