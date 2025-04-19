package main

import(
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
    n := len(nums)
    xorr := n
    for i := 0; i < n; i++ {
        xorr ^= i ^ nums[i]
    }
    return xorr
}
