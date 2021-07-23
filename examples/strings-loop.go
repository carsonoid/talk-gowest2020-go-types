package main

import "fmt"

func main() {
	const s = `西に行く`

	for i := 0; i < len(s); i++ {
		fmt.Println(
			s[i], ":", string(s[i]),
		)
	}
}
