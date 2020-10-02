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
	time.Sleep(time.Second * sleepSeconds)
	fmt.Println("done")
}
