package main

import "fmt"

func main() {
	arr := [3]bool{true, false, true}
	arrCopy := arr

	// Now change the copy
	arrCopy[1] = true

	fmt.Println("arr", arr)
	fmt.Println("arrCopy", arrCopy)
}
