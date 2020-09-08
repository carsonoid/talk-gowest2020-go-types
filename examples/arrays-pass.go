package main

import "fmt"

func main() {
	var bools [5]bool
	fmt.Println(bools)

	changeBoolsToTrue(bools)
	fmt.Println(bools)
}

func changeBoolsToTrue(bools [5]bool) {
	for i := range bools {
		bools[i] = true
	}
}

// END OMIT
