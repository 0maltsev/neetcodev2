package main

import "fmt"

func main() {
	temperatures := []int{73,74,75,71,69,72,76,73}
	fmt.Println(dailyTemperatures(temperatures))
}

type StackObject struct {
	Value int
	Index int
}

func dailyTemperatures(temperatures []int) []int {
	result := []int{}
	stack := []StackObject{}
	tempLen := len(temperatures)
	initResultArray(&result, tempLen)
	result[tempLen-1] = 0
	stack = append(stack, StackObject{Value: temperatures[tempLen-1], Index: tempLen-1})
	for i:= tempLen-2; i>=0; i-- {
		for len(stack) > 0 && temperatures[i] >= stack[len(stack)-1].Value{
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1].Index-i
		}
		stack = append(stack, StackObject{Value: temperatures[i], Index: i})
	}
	return result
}

func initResultArray(result *[]int, l int){
	for i:=0; i < l; i++{
		*result = append(*result, 0)
	}
}