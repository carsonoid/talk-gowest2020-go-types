package main

import (
	"fmt"
	"time"
)

func main() {
	const sleepSeconds time.Duration = 2

	fmt.Print("sleeping for ", sleepSeconds, " seconds...")
	time.Sleep(time.Second * sleepSeconds) // HL
	fmt.Println("done")
}
