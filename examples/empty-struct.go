package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var empty struct{}
	fmt.Println(unsafe.Sizeof(empty))
}
