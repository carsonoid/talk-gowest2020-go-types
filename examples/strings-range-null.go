package main

import "fmt"

func main() {
	raw := []byte(string(`西に`))
	raw = append(raw, 0b0) // add NULL byte in the middle
	raw = append(raw, []byte(string(`行く`))...)

	s := string(raw)

	for _, x := range s {
		fmt.Println(
			x, ":", string(x),
		)
	}
}
