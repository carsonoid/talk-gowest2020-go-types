package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ui8 uint8
	fmt.Println(unsafe.Sizeof(ui8)*8, "bits")

	var ui16 uint16
	fmt.Println(unsafe.Sizeof(ui16)*8, "bits")

	var ui32 uint32
	fmt.Println(unsafe.Sizeof(ui32)*8, "bits")

	var ui64 uint64
	fmt.Println(unsafe.Sizeof(ui64)*8, "bits")
}
