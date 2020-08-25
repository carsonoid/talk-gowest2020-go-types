package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var data struct {
		x bool   // 1 bytes
		y uint64 // 8 bytes
	}
	fmt.Println(
		unsafe.Sizeof(data.x),
		"+",
		unsafe.Sizeof(data.y),
		"=",
		unsafe.Sizeof(data),
	)
}
