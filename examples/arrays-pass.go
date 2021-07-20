package main

import "fmt"

func main() {
	var arr [3]bool
	changeToTrue(arr)
	fmt.Println(arr)
}

func changeToTrue(arr [3]bool) {
	for i := range arr {
		arr[i] = true
	}
}

// END OMIT
