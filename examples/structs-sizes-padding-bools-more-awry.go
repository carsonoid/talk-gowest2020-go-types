package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var data struct {
		x1 bool // 1 bytes
		// [7]byte
		y  uint64 // 8 bytes
		x2 bool   // 1 bytes
		// [7]byte
	}
	fmt.Println(
		unsafe.Sizeof(data.x1),
		"+",
		unsafe.Sizeof(data.x2),
		"+",
		unsafe.Sizeof(data.y),
		"=",
		unsafe.Sizeof(data),
	)
}
