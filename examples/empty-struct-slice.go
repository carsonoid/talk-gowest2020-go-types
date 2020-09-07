package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// use make to create a slice of len 100
	structSlice := make([]struct{}, 100)
	fmt.Println(unsafe.Sizeof(structSlice))
}
