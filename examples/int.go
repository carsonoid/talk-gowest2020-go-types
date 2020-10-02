package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	fmt.Println(unsafe.Sizeof(i)*8, "bits")

	var u uint
	fmt.Println(unsafe.Sizeof(u)*8, "bits")
}
