package main

import "fmt"

func main() {
	var (
		x  = 'x'
		y  = '😊'
		ni = '日'
	)

	fmt.Print(
		[]byte(string(x)), "\n",
		[]byte(string(y)), "\n",
		[]byte(string(ni)), "\n",
	)

}
