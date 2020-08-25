package main

import (
	"fmt"
	"reflect"
)

func main() {
	var ui8 uint8
	sizeInBits(ui8)

	var ui16 uint16
	sizeInBits(ui16)

	var ui32 uint32
	sizeInBits(ui32)

	var ui64 uint64
	sizeInBits(ui64)
}

func sizeInBits(i interface{}) {
	// use reflect to get size in bytes
	s := reflect.TypeOf(i).Size()

	// print bytes * 8 to get bits
	fmt.Println(s * 8)
}
