package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// empty struct with only empty struct children has zero size
	var nestedStruct struct {
		a struct{}
		b struct{}
	}
	fmt.Println(unsafe.Sizeof(nestedStruct))
}
