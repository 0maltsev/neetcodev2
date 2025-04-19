package main

import(
	"fmt"
)

func main() {
	fmt.Println(reverseBits(00000000000000000000000000010101))
}

func reverseBits(n uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		bit := (n >> i) & 1
		res |= (bit << (31 - i))
	}
	return res
}
