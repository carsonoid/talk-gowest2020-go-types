package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// array has zero size
	var structs [100]struct{}
	fmt.Println(unsafe.Sizeof(structs))

	// bool array for comparison
	var bools [100]struct{}
	fmt.Println(unsafe.Sizeof(bools))
}
