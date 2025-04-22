package main

import()

func main() {

}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	const START, END = 0, 1
	result := [][]int{}
	firstPointer, secondPointer := 0, 0

	for firstPointer < len(firstList) && secondPointer < len(secondList) {
		start1, end1 := firstList[firstPointer][START], firstList[firstPointer][END]
		start2, end2 := secondList[secondPointer][START], secondList[secondPointer][END]

		if end1 < start2 {
			firstPointer++ // period1 is before period2
		} else if end2 < start1 {
			secondPointer++ // period2 is before period1
		} else { // they intersect, move the pointer of the period that ends first
			result = append(result, []int{max(start1, start2), min(end1, end2)})
			if end1 < end2 {
				firstPointer++
			} else {
				secondPointer++
			}
		}
	}
	return result
}

