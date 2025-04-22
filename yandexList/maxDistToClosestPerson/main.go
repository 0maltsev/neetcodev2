package main

import()

func main() {

}

func maxDistToClosest(seats []int) int {
	var leftGap, rightGap, maxGap int
	start, end := -1, -1
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			leftGap = i
			start = i
			break
		}
	}
	for i := len(seats) - 1; i >= 0; i-- {
		if seats[i] == 1 {
			rightGap = len(seats) - 1 - i
			end = i
			break
		}
	}
	for prev, i := start, start+1; i <= end; i++ {
		if seats[i] == 1 {
			maxGap = max(maxGap, i-prev)
			prev = i
		}
	}
	return max(leftGap, rightGap, maxGap/2)
}
