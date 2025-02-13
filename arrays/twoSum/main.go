package main

import "fmt"

func main(){
	nums := []int{3,2,4}
	target := 6

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := arrayToHashMap(nums)
	for index, element := range(nums) {
		m[element] = m[element][1:]
		if len(m[target-element]) != 0 {
			return []int{index, m[target-element][0]}
		}
	}
	return []int{}
}

func arrayToHashMap(array []int) map[int][]int {
	m := make(map[int][]int)
	for index, element := range array {
		m[element] = append(m[element], index)
	}
	return m
}
