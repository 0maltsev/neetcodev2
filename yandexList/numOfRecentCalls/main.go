package main

import()

func main() {}

type RecentCounter struct {
    requests []int
}

func Constructor() RecentCounter {
    return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
    // Add the current timestamp to the queue
    this.requests = append(this.requests, t)
    
    // Remove timestamps that are outside the 3000 milliseconds window
    for this.requests[0] < t-3000 {
        this.requests = this.requests[1:]
    }
    
    // Return the number of requests within the window
    return len(this.requests)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
