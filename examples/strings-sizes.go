package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "I am a string"
	fmt.Println("string size:        ", unsafe.Sizeof(s))

	s = "I am a longer string"
	fmt.Println("longer string size: ", unsafe.Sizeof(s))
}
