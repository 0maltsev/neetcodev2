package main

import(
	"fmt"
)

func main() {

}

func summaryRanges(nums []int) (result []string) {
	l := len(nums)

	if l == 0 {
		return
	}

	s := 0
	for i := 1; i <= l; i++ {
		if i == l || nums[i] != nums[i-1]+1 {
			if nums[s] == nums[i-1] {
				result = append(result, fmt.Sprintf("%d", nums[s]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[s], nums[i-1]))
			}
			if i < l {
				s = i
			}
		}
	}

	return result
}
