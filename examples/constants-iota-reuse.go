package main

import "fmt"

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0  (iota == 0)
	bit1, mask1                          // bit1 == 2, mask1 == 1  (iota == 1)
	_, _                                 // unused
	bit3, mask3                          // bit3 == 8, mask3 == 7  (iota == 3)
)

// END OMIT

func main() {
	fmt.Println(bit0, mask0)
	fmt.Println(bit1, mask1)
	fmt.Println(bit3, mask3)
}
