package main

import(
	"fmt"
)

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000010111))
}

func hammingWeight(n int) int {
	res := 0
	for n != 0 {
		n &= n - 1
        res++
	}
	return res
}
