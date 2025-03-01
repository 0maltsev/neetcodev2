package main

import "fmt"

func main() {
	s := "(){}}{"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := []string{}
	m := make(map[string]string)
	m["{"] = "}"
	m["("] = ")"
	m["["] = "]"
	for _, r := range s {
		if m[string(r)] == "" {
			if len(stack) == 0 {
				return false
			}
			if m[stack[len(stack)-1]] == string(r) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, string(r))
		}
	}
	return len(stack) == 0
}