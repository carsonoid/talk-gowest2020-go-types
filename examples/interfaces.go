package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// BEGIN OMIT
func printSize(msg string, i interface{}) {
	fmt.Printf("\nsizes in bytes for %s '%v':\n", msg, i)
	fmt.Println("interface:  ", unsafe.Sizeof(i))
	fmt.Println("type: ", reflect.TypeOf(i).Size())
}

func main() {
	printSize("bool", true)
	printSize("untyped const", 1)
	printSize("int32 cast const", int32(1))
	printSize("int64 cast const", int64(1))
	printSize("string", "xxxx")
}

// END OMIT
