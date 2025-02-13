package main

import "fmt"

func main() {
	nums := []int{2}
	result := containsDuplicate(nums)
	fmt.Println(result)
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, element := range (nums) {
		m[element] = m[element] + 1
		if m[element] > 1 {
			return true
		}
	}
	return false
}