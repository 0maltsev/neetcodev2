package main

import "fmt"

func main() {
	// nums := []int{1,2,3,4,5}
	nums := []int{2,3,4,5,1}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
    res := nums[0]
    l, r := 0, len(nums)-1

    for l <= r {
        if nums[l] < nums[r] {
            if nums[l] < res {
                res = nums[l]
            }
            break
        }

        m := l + (r-l)/2
        if nums[m] < res {
            res = nums[m]
        }

        if nums[m] >= nums[l] {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return res
}