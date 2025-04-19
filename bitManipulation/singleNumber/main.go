package main

import(
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
