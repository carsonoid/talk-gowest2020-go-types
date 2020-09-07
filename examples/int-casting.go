package main

import "fmt"

func main() {
	var i32 int32 = 1002
	var i64 int64 = 1002

	fmt.Println(i64 == int64(i32))
	fmt.Println(i32 == int32(i64))
}
