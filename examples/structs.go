package main

import "fmt"

// BEGIN OMIT
type nothin struct{}

type data struct {
	x int
	y int64
}

func main() {
	fmt.Println(nothin{})
	fmt.Println(data{1, 2})
}

// END OMIT
