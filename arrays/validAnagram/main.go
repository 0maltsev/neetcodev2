package main

import "fmt"

func main() {
	s := "assa"
	t := "saas"
	result := isAnagram(s, t)
	fmt.Println(result)
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m_s, t_s := convertStringToHasmap(s, t)
	for letter, amount := range(m_s) {
		if amount != t_s[letter] {
			return false
		}
	}
	return true
}

func convertStringToHasmap(s string, t string) (map[rune]int, map[rune]int) {
	t_s := make(map[rune]int)
	m_s := make(map[rune]int)
	for i, letter := range(s){
		m_s[letter] += 1
		t_s[rune(t[i])] += 1 
	}
	return m_s, t_s
}