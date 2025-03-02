package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
    operationList := []string{"+", "-", "/", "*"}
	stack := []string{}

	for i := range tokens {
		if !contains(operationList, tokens[i]) {
			stack = append(stack, tokens[i])
		} else {
			stack = operationOnStack(stack, tokens[i])
		}
	}
	result, _ := strconv.ParseInt(stack[0], 10, 32)
	return int(result)
}

func operationOnStack(stack []string, operation string) []string{
	stack[len(stack)-2] = strconv.Itoa(resultByStringOperator(operation, stack[len(stack)-2], stack[len(stack)-1]))
	stack = stack[:len(stack)-1]
	return stack
}

func resultByStringOperator(operator string, a string, b string) int {
	aInt, _ := strconv.ParseInt(a, 10, 32)
	bInt, err := strconv.ParseInt(b, 10, 32)
	if err == nil {
		switch operator {
			case "+":
				return sumOperation(int(aInt), int(bInt))
			case "-":
				return minusOperation(int(aInt), int(bInt))
			case "/":
				return devideOperation(int(aInt), int(bInt))
			case "*":
				return multOperation(int(aInt), int(bInt))
		}
	}
	return 0
}

func sumOperation(a,b int) int {
	return a+b
}

func minusOperation(a,b int) int {
	return a-b
}

func multOperation(a,b int) int {
	return a*b
}

func devideOperation(a,b int) int {
	return a/b
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}