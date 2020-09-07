package main

import "fmt"

// BEGIN OMIT
const (
	a            = 0 // iota
	b int64      = 1 // iota
	c float64    = 2 // iota
	d complex128 = 3 // iota
)

func main() {
	// Print value and type of all constants
	fmt.Printf("%[1]v:%[1]T\n", a)
	fmt.Printf("%[1]v:%[1]T\n", b)
	fmt.Printf("%[1]v:%[1]T\n", c)
	fmt.Printf("%[1]v:%[1]T\n", d)
}

// END OMIT
