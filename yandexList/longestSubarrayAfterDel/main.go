package main

func main() {

}

func longestSubarray(nums []int) int {
    zeros_count := 0 
    start := 0 
    max_length := 0
    for end:= 0 ; end < len(nums) ; end ++ {
        if nums[end] == 0 {
            zeros_count ++
        }
        
        for zeros_count > 1 {
            if nums[start] == 0 {
                zeros_count --
            }
            start ++
        }
        max_length = max(max_length , end -start)
    }

    return max_length 
}