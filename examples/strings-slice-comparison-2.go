package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("string sizes:  ")
	var str string = "Hi"
	fmt.Println("header: ", unsafe.Sizeof(str))
	fmt.Println("len:    ", len(str))
	fmt.Println("total:  ", int(unsafe.Sizeof(str))+len(str))

	fmt.Println()

	fmt.Println("byte slice sizes:  ")
	slice := []byte(str)
	fmt.Println("header: ", unsafe.Sizeof(slice))
	fmt.Println("cap:    ", cap(slice)) // you have to use cap for slices
	fmt.Println("total:  ", int(unsafe.Sizeof(slice))+cap(slice))
}
