package main

import (
	"fmt"
	"time"
)

func main() {
	// BEGIN OMIT
	const sleepSeconds = 2
	// END OMIT

	fmt.Print("sleeping for ", sleepSeconds, " seconds...")
	// invalid code do demonstrate the idea beind untyped constants OMIT
	time.Sleep(time.Second * sleepSeconds.(time.Duration)) // HL
	fmt.Println("done")
}
