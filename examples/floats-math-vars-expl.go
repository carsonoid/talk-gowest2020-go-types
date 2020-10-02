package main

import "fmt"

func main() {
	var (
		a = 0.1
		b = 0.2
		c = 0.3
	)

	fmt.Printf("a:   %.20f\n", a)
	fmt.Printf("b:   %.20f\n", b)
	fmt.Printf("a+b: %.20f\n", a+b)
	fmt.Printf("c:   %.20f\n", c)
}
