package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i8 int8
	fmt.Println(unsafe.Sizeof(i8)*8, "bits")

	var i16 int16
	fmt.Println(unsafe.Sizeof(i16)*8, "bits")

	var i32 int32
	fmt.Println(unsafe.Sizeof(i32)*8, "bits")

	var i64 int64
	fmt.Println(unsafe.Sizeof(i64)*8, "bits")
}
