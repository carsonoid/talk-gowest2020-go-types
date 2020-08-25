package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i8 int8
	sizeInBits(i8)

	var i16 int16
	sizeInBits(i16)

	var i32 int32
	sizeInBits(i32)

	var i64 int64
	sizeInBits(i64)
}

func sizeInBits(i interface{}) {
	// use reflect to get size in bytes
	s := reflect.TypeOf(i).Size()

	// print bytes * 8 to get bits
	fmt.Println(s * 8)
}
