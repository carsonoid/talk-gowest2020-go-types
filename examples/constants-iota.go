package main

import "fmt"

const (
	Created = iota
	Starting
	Started
	Done
)

func main() {
	fmt.Println(Created, Starting, Started, Done)
}
