package main

func main() {
	var data struct {
		x uint32 // 4 bytes
		_ [4]byte
		y uint64 // 8 bytes
	}
}
