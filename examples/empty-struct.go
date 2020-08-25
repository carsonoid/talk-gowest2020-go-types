package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// empty struct has zero size
	var empty struct{}
	fmt.Println(unsafe.Sizeof(empty))
}
