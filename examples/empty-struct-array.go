package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var structArray [100]struct{}
	fmt.Println(unsafe.Sizeof(structArray))
}
