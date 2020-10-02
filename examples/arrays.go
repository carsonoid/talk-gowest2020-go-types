package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var bools [50]bool
	fmt.Println("size is", unsafe.Sizeof(bools), "bytes")

	fmt.Println("index 0  is:", bools[0])

	bools[42] = true
	fmt.Println("index 42 is:", bools[42])
}
