package main

import "fmt"

func main() {
	nums := []int{100,4,200,1,3,2}
	fmt.Print(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	var result int

	m := make(map[int]int)
    for _, element := range(nums) {
		m[element]++
	}

	for element := range m {
		if m[element-1] == 0 {
			counter := 1
			for {
				if m[element+counter] != 0 {
					counter ++
				} else {
					break
				}
			}
		if counter > result {
			result = counter
			}
		}
	}
	
	return result
}