package main

import (
	"fmt"
	"unsafe"
)

type position struct {
	x, y int8
}

func main() {
	fmt.Println("position struct sizes")

	// empty vector struct is two bytes
	var p position
	fmt.Println("single: ", unsafe.Sizeof(p))

	// array of positions is len * 2 bytes, or 200 bytes
	var positionArray [100]position
	fmt.Println("len 100 array: ", unsafe.Sizeof(positionArray))

	// slize of position slice is always slice header length
	// note that the underlying array for this slice still exists
	// and still uses 200 bytes for a capacity of 100
	positionSlice := make([]position, 100)
	fmt.Println("len 100 slice: ", unsafe.Sizeof(positionSlice))
}
