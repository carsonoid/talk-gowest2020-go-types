package main

import (
	"fmt"
	"reflect"
)

func main() {
	// by explicitly converting to a rune slice
	// you can get indexes in a more natural way
	// NOTE: This is still not 100% reliable!
	s := "I ám a string"
	x := []rune(s)[3] // HL
	fmt.Println("at index 3: ", x)
	fmt.Println("type is:    ", reflect.TypeOf(x))
	fmt.Println("as string:  ", string(x))
}
