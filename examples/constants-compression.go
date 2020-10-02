package main

import "fmt"

const (
	Created, Starting, Started, Done = iota, iota, iota, iota
)

func main() {
	fmt.Println(Created, Starting, Started, Done)
}
