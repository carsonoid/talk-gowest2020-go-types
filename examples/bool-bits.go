package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b bool
	fmt.Println(unsafe.Sizeof(b)*8, "bits")
}

// Partial copy of https://golang.org/src/unsafe/unsafe.go?s=8189:8225#L178
// START OMIT
// Sizeof takes an expression x of any type and returns the size in bytes
// of a hypothetical variable v as if v was declared via var v = x.
// The size does not include any memory possibly referenced by x.
// For instance, if x is a slice, Sizeof returns the size of the slice
// descriptor, not the size of the memory referenced by the slice.
// The return value of Sizeof is a Go constant.
func Sizeof(x ArbitraryType) uintptr {
	// END OMIT
	return 0
}

type ArbitraryType string
