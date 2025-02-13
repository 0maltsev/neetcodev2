package main

import "fmt"

func main(){
	nums := []int{2,7,11,15}
	target := 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, element := range(nums) {
		if m[target - element] != 0 {
			return []int{m[target - element]-1, index}
		}
		m[element] = index + 1
	}
	return []int{}
}
