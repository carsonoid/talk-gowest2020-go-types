package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int
	sizeInBits(i)
}

func sizeInBits(i interface{}) {
	// use reflect to get size in bytes
	s := reflect.TypeOf(i).Size()

	// print bytes * 8 to get bits
	fmt.Println(s * 8)
}
