package main

import (
	"math"
)

func main() {

}

type MinStack struct {
    min   int
    stack []int
}

func Constructor() MinStack {
    return MinStack{
        min:   math.MaxInt64,
        stack: []int{},
    }
}

func (s *MinStack) Push(x int) {
    if len(s.stack) == 0 {
        s.stack = append(s.stack, 0)
        s.min = x
    } else {
        s.stack = append(s.stack, x - s.min)
        if x < s.min {
            s.min = x
        }
    }
}

func (s *MinStack) Pop() {
    if len(s.stack) == 0 {
        return
    }
    pop := s.stack[len(s.stack)-1]
    s.stack = s.stack[:len(s.stack)-1]
    if pop < 0 {
        s.min = s.min - pop
    }
}

func (s *MinStack) Top() int {
    top := s.stack[len(s.stack)-1]
    if top > 0 {
        return top + s.min
    }
    return s.min
}

func (s *MinStack) GetMin() int {
    return s.min
}