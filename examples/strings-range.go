package main

import "fmt"

func main() {
	const s = `西に行く`

	for _, x := range s {
		fmt.Println(
			x, ":", string(x),
		)
	}
}
