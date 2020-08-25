package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var data struct {
		child struct {
			a uint64 // 8 bytes
			b uint64 // 8 bytes
		}
		x uint64 // 8 bytes
		y uint64 // 8 bytes
	}

	fmt.Println(
		unsafe.Sizeof(data.child.a),
		"+",
		unsafe.Sizeof(data.child.b),
		"+",
		unsafe.Sizeof(data.x),
		"+",
		unsafe.Sizeof(data.y),
		"=",
		unsafe.Sizeof(data),
	)
}
