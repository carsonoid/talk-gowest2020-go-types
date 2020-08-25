package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// strut slice with members has only slice header size
	structsSlice := make([]struct{}, 100)
	fmt.Println(unsafe.Sizeof(structsSlice))
	arraySize := cap(structsSlice) * int(unsafe.Sizeof(struct{}{}))
	fmt.Println("array size: ", arraySize)

	// add and size doesn't change
	structsSlice = append(structsSlice, struct{}{})
	fmt.Println(unsafe.Sizeof(structsSlice))
	arraySize = cap(structsSlice) * int(unsafe.Sizeof(struct{}{}))
	fmt.Println("array size: ", arraySize)
}
