package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("sleeping for 2 seconds...")
	time.Sleep(time.Second * 2)
	fmt.Println("done")
}
