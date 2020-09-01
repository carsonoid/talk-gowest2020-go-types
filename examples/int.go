package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	fmt.Println(unsafe.Sizeof(i)*8, "bits")
}
