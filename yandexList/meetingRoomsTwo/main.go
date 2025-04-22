package main

import("sort")

func main() {}

type Interval struct {
	start int
	end   int
}

func minMeetingRooms(intervals []Interval) int {
    start := make([]int, len(intervals))
    end := make([]int, len(intervals))
    
    for i, interval := range intervals {
        start[i] = interval.start
        end[i] = interval.end
    }
    
    sort.Ints(start)
    sort.Ints(end)
    
    res, count := 0, 0
    s, e := 0, 0
    
    for s < len(intervals) {
        if start[s] < end[e] {
            s++
            count++
        } else {
            e++
            count--
        }
        if count > res {
            res = count
        }
    }
    
    return res
}
