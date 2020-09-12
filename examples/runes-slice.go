package main

import "fmt"

func main() {
	var (
		str = "Hi"
		x   = 'x'
		y   = '😊'
		ni  = '日'
	)

	fmt.Print(
		// string to []byte
		[]byte(str), "\n",
		// or with runes by converting to strings first
		[]byte(string(x)), "\n",
		[]byte(string(y)), "\n",
		[]byte(string(ni)), "\n",
	)

}
