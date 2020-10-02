package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var data struct {
		x uint32 // 4 bytes
		y uint32 // 4 bytes
	}
	fmt.Println(
		unsafe.Sizeof(data.x),
		"+",
		unsafe.Sizeof(data.y),
		"=",
		unsafe.Sizeof(data),
	)
}
