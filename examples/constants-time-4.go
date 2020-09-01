package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	sleepSeconds := 2

	if len(os.Args) > 1 {
		sleepSeconds, _ = strconv.Atoi(os.Args[1])
	}

	fmt.Print("sleeping for ", sleepSeconds, " seconds...")
	time.Sleep(time.Second * time.Duration(sleepSeconds)) // HL
	fmt.Println("done")
}
