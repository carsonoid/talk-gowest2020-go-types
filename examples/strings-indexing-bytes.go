package main

import (
	"fmt"
	"reflect"
)

func main() {
	buf := []byte("I á")
	buf = append(buf, 0x00) // write nul byte
	buf = append(buf, 'm')  // write ascii 'a' byte
	s := string(buf) + " a string"
	fmt.Printf("s is : '%s'\n", s)
	x := []rune(s)[3]
	fmt.Println("at index 3: ", x)
	fmt.Println("type is:    ", reflect.TypeOf(x))
	fmt.Println("as string:  ", string(x))
}
