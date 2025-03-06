package main

import "fmt"

func main() {
	heights := []int{1,1}
	fmt.Println(maxArea(heights))
}

func maxArea(height []int) int {
	var max int
	i := 0
	j := len(height)-1
	for i < j {
		newMax := getArea(i, j, height)
		if  newMax > max {
			max = newMax
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func getArea(i, j int, arr []int) int {
	if arr[i] < arr[j] {
		return arr[i]*(j-i)
	}
	return arr[j]*(j-i)
}
