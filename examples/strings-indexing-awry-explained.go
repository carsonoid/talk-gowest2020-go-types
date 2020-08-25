package main

import (
	"fmt"
)

func main() {
	s := "I am a string"
	fmt.Println([]byte(s))
	x := s[3]
	fmt.Println("at index 3: ", x)

	fmt.Println()

	s = "I ám a string"
	fmt.Println([]byte(s))
	x = s[3]
	fmt.Println("at index 3: ", x)
}
