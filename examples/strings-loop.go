package main

import "fmt"

func main() {
	// s == "go west" in Kanji
	const s = `西に行く`

	for i := 0; i < len(s); i++ {
		fmt.Println(
			s[i], ":", string(s[i]),
		)
	}
}
