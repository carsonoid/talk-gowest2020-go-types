package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var f32 float32 = 1.002
	fmt.Println(unsafe.Sizeof(f32)*8, "bits")

	var f64 float64 = 1.002
	fmt.Println(unsafe.Sizeof(f64)*8, "bits")
}
