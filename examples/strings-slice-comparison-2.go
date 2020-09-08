package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string = "Hi"
	fmt.Println("string bytes: ", unsafe.Sizeof(str))
	fmt.Println("string size:  ", len(str))

	slice := []byte(str)
	fmt.Println("slice bytes:  ", unsafe.Sizeof(slice))
	fmt.Println("slice size:   ", len(slice))
}
