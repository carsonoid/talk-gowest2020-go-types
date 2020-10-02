package main

import "fmt"

// BEGIN OMIT
const (
	a            = iota
	b int64      = iota
	c float64    = iota
	d complex128 = iota
)

// END OMIT

func main() {
	// Print value and type of all constants
	fmt.Printf("%[1]v:%[1]T\n", a)
	fmt.Printf("%[1]v:%[1]T\n", b)
	fmt.Printf("%[1]v:%[1]T\n", c)
	fmt.Printf("%[1]v:%[1]T\n", d)
}
