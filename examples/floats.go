package main

import "fmt"

func main() {
	const fconst = 1.002
	var f64 float64 = 1.002
	var f32 float32 = 1.002

	fmt.Println(fconst, "==", f64, ":", fconst == f64)
	fmt.Println(fconst, "==", f32, ":", fconst == f32)

	fmt.Println(f32, "==", f64, ":", f32 == float32(f64))
	fmt.Println(f64, "==", f32, ":", f64 == float64(f32))
}
