package main

import()

func main() {}

func findMaxConsecutiveOnes(nums []int) int {
	left, counter := 0, 0
	for _, x := range nums {
		counter += x ^ 1
		if counter > 1 {
			counter -= nums[left] ^ 1
			left++
		}
	}
	return len(nums) - left
}
