package main

import(
	"fmt"
)

func main() {
	nums := []int{-1,0,3,5,9,12}
	target := 9
	fmt.Println(search(nums, target))
}

func binarySearch(l, r int, nums []int, target int) int {
    if l > r {
        return -1
    }
    m := l + (r-l)/2

    if nums[m] == target {
        return m
    }
    if nums[m] < target {
        return binarySearch(m+1, r, nums, target)
    }
    return binarySearch(l, m-1, nums, target)
}

func search(nums []int, target int) int {
    return binarySearch(0, len(nums)-1, nums, target)
}