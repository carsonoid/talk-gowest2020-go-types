package main

import (
	"fmt"
	"unsafe"
)

func main() {
	bools := []bool{}
	fmt.Println("bools size: ", unsafe.Sizeof(bools))
	arraySize := cap(bools) * int(unsafe.Sizeof(true))
	fmt.Println("array size: ", arraySize)

	fmt.Println()

	ints := make([]int8, 100)
	fmt.Println("ints size:  ", unsafe.Sizeof(ints))
	arraySize = cap(ints) * int(unsafe.Sizeof(int(1)))
	fmt.Println("array size: ", arraySize)
}
