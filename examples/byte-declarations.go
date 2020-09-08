package main

import "fmt"

var (
	b1 byte = 0         // integer literal
	b2 byte = 0xa       // hex notation
	b3 byte = 'd'       // rune notation
	b4 byte = 1_01      // integer literal w/underscore
	b5 byte = 0b1100110 // binary notation
)

func main() {
	fmt.Println(b1, b2, b3, b4, b5)
}
