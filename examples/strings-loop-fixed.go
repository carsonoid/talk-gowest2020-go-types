package main

import "fmt"

func main() {
	const s = `西に行く`

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Println(
			runes[i], ":", string(runes[i]),
		)
	}
}
