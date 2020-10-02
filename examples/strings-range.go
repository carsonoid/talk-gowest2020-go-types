package main

import "fmt"

func main() {
	// s == "go west" in Kanji
	const s = `西に行く`

	for _, x := range s {
		fmt.Println(
			x, ":", string(x),
		)
	}
}
