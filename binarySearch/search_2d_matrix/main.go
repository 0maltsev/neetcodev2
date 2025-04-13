package main

import "fmt"

func main() {
	matrix := [][]int{
		{1,3,5,7},
		{10,11,16,20},
		{23,30,34,60},
	}
	target := 13
	fmt.Println(searchMatrix(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	searchRowNum := 0
    for rowNum, row := range matrix {
		if row[len(row)-1] >= target {
			searchRowNum = rowNum
			break
		}
	}
	return search(matrix[searchRowNum], target)
}

func binarySearch(l, r int, nums []int, target int) bool {
    if l > r {
        return false
    }
    m := l + (r-l)/2

    if nums[m] == target {
        return true
    }
    if nums[m] < target {
        return binarySearch(m+1, r, nums, target)
    }
    return binarySearch(l, m-1, nums, target)
}

func search(nums []int, target int) bool {
    return binarySearch(0, len(nums)-1, nums, target)
}