package main

import("sort")

func main() {}

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    output := [][]int{intervals[0]}

    for _, interval := range intervals[1:] {
        start, end := interval[0], interval[1]
        lastEnd := output[len(output)-1][1]

        if start <= lastEnd {
            output[len(output)-1][1] = max(lastEnd, end)
        } else {
            output = append(output, []int{start, end})
        }
    }
    return output
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
