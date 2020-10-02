package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// BEGIN OMIT
	sleepSeconds := 2
	// END OMIT

	if len(os.Args) > 1 {
		sleepSeconds, _ = strconv.Atoi(os.Args[1])
	}

	fmt.Print("sleeping for ", sleepSeconds, " seconds...")
	time.Sleep(time.Second * sleepSeconds)
	fmt.Println("done")
}
