package main

import (
	"fmt"
)

func main() {
	var x *int
	var y interface{}

	fmt.Println("x == nil:", x == nil)
	fmt.Println("y == nil:", x == nil)
	fmt.Println("x == y:  ", x == y)
}
