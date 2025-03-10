package main

import "fmt"

func main() {
	target := 9
	numbers := []int{2,7,11,15}
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	result := []int{}
	for i, j:= 0, len(numbers)-1; i < j; {
		sum := numbers[i] + numbers[j]
		if sum > target {
			j = j-1
		} else if sum < target {
			i = i+1
		} else {
			result = append(result, i+1)
			result = append(result, j+1)
			return result
		}
	}
	return result
}
