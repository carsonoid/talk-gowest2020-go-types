package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string = "Hi"
	fmt.Println("string bytes: ", unsafe.Sizeof(str))

	slice := []byte(str)
	fmt.Println("slice bytes:  ", unsafe.Sizeof(slice))
}
