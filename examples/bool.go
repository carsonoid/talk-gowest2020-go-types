package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b bool
	fmt.Println(unsafe.Sizeof(b))
}
