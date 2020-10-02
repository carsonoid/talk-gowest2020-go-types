package main

import "fmt"

func main() {
	var f32 float32 = 1.002
	var f64 float64 = 1.002

	fmt.Println(f64 == float64(f32))
	fmt.Println(f32 == float32(f64))
}
