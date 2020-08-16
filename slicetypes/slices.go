package main

import "fmt"

func main() {
	// bools := []bool{true, true, false, false}
	// printBools(bools)

	// rangeSlices()

	sliceCopies()

	// stringSliceConversion()
}

// Go Passes everything by copy. How many values are copied
// when we run this function?
// one value - A slice pointer
// three values - All list elements
// four values - All list elements
func printBools(bools []bool) {
	for i, b := range bools {
		fmt.Println("index", i, "is", b)
	}
}

// why can you range over a nil slice?
func rangeSlices() {
	// Yep!
	var bools []bool
	for _, b := range bools {
		fmt.Println(b)
	}

	// And you can append to one too!
	// very handy with maps of slices
	bools = append(bools, true)
}

// Do slice copies cause allocations?
func sliceCopies() {
	slice := []byte{0x1, 0x2, 0x3}
	fmt.Println(slice[0], &slice[0])

	newSlice := slice
	fmt.Println(newSlice[0], &newSlice[0])

	// modify a member
	slice[0] = 0x4
	fmt.Println(slice[0], &slice[0])
	fmt.Println(newSlice[0], &newSlice[0])
}

// Do slice conversions caus allocations?
// i.e. is new data created when you `s := string(slice)`?
func stringSliceConversion() {
	slice := []byte{0x1, 0x2, 0x3}
	fmt.Println(&slice[0])

	str := string(slice)

	newSlice := []byte(str)
	fmt.Println(&newSlice[0])

	// modify a member
	slice[0] = 0x4
	fmt.Println(slice[0], &slice[0])
	fmt.Println(newSlice[0], &newSlice[0])
}
