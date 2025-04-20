package main

func main() {

}

func subarraySum(nums []int, k int) int {
	m := make(map[int]int) // prefix_sum -> count
	m[0] = 1

	prefixSum := 0
	result := 0
	for _, num := range nums {
		prefixSum += num
		if count, exists := m[prefixSum-k]; exists {
			result += count
		}
		m[prefixSum]++
	}
	return result
}
