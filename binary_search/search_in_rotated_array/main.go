package main

import "fmt"

func main() {
	nums := []int{4,5,6,7,0,1,2}
	target := 9
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
    l, r := 0, len(nums)-1

    for l <= r {
        mid := (l + r) / 2
        if target == nums[mid] {
            return mid
        }

        if nums[l] <= nums[mid] {
            if target > nums[mid] || target < nums[l] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        } else {
            if target < nums[mid] || target > nums[r] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        }
    }
    return -1
}