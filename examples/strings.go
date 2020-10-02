package main

import "fmt"

const goWestEN = `GoWest`
const goWestJP = `西に行く`

func main() {
	fmt.Println("-- ASCII --")
	examineString(goWestEN)

	fmt.Println("-- Non-ASCII --")
	examineString(goWestJP)
}

func examineString(s string) {
	fmt.Println("Whole string: ", s)

	fmt.Println("Index 0: ", s[0])

	fmt.Println("for loop:")
	for i := 0; i < len(s); i++ {
		b := s[i]
		// the default Println behavor is to print the type (byte)
		fmt.Println("rune: ", b)

		// we have to cast to string to get our character
		fmt.Printf("quote: %s\n", string(b))
	}

	fmt.Println("for range loop:")
	for _, r := range s {
		// the default Println behavor is to print the type (byte slice)
		fmt.Println("rune: ", r)

		// we have to cast to string to get our character
		fmt.Printf("quote: %s\n", string(r))
	}
}
