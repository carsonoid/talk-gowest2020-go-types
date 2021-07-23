package main

import (
	"fmt"
)

func main() {
	str := "Hi"

	slice := []byte(str)

	slice[1] = 'm'
	slice = append(slice, 109, 109, 109, 240, 159, 164, 148)

	// print proof
	fmt.Println(str)
	fmt.Println(string(slice))
}
