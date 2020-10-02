// This file is in it's own folder because that is a requirement for
// go generate to work.
package main

import "fmt"

// BEGIN OMIT
//go:generate stringer -type=State // HL
type State int

const (
	Created State = iota
	Starting
	Started
	Done
)

func main() {
	fmt.Println(Created, Starting, Started, Done)
}

// END OMIT
