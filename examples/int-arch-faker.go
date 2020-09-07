// This file fakes the output of
// GOARCH=386 go run examples/int.go

package main

import (
	"fmt"
)

// BEGIN OMIT
// command: GOARCH=386 go run examples/int.go
// END OMIT

func main() {
	fmt.Println("32 bits")
	fmt.Println("32 bits")
}
