package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var nestedStruct struct {
		a struct{}
		b struct{}
	}
	fmt.Println(unsafe.Sizeof(nestedStruct))
}
