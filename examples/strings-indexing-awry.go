package main

import (
	"fmt"
	"reflect"
)

func main() {
	// strings with non-ascii bytes will mess up indexing
	s := "I ám a string"
	x := s[3]
	fmt.Println("at index 3: ", x)
	fmt.Println("type is:    ", reflect.TypeOf(x))
	fmt.Println("as string:  ", string(x))
}
