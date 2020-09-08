package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	fmt.Print("Sleeping...")
	go func() {
		time.Sleep(time.Second)
		done <- struct{}{}
	}()

	<-done
	fmt.Println("done")
}
