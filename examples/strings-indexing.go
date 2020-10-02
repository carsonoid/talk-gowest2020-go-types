package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ascii only strings can be safely indexed as normal
	s := "I am a string"
	x := s[3]
	fmt.Println("at index 3: ", x)
	fmt.Println("type is:    ", reflect.TypeOf(x))
	fmt.Println("as string:  ", string(x))
}
