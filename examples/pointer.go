package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b bool
	fmt.Println("bool size:", unsafe.Sizeof(b))
	fmt.Println("bool pointer size:", unsafe.Sizeof(&b))

	var i64 int64
	fmt.Println("int64 size:", unsafe.Sizeof(i64))
	fmt.Println("int64 pointer size:", unsafe.Sizeof(&i64))

	s := []bool{}
	fmt.Println("slice size:", unsafe.Sizeof(s))
	fmt.Println("slice pointer size:", unsafe.Sizeof(&s))
}
