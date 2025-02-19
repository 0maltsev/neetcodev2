package main

import (
	"fmt"
)

func main() {
	strs := []string{"act","pots","tops","cat","stop","hat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := make(map[[26]int][]string)
	
	for _, str := range strs {
		count := [26]int{}
		for _, char := range(str) {
			count[char - 'a']++
		}
		m[count] = append(m[count], str)
	}
	for _, element := range m {
		res = append(res, element)
	}
	return res
}
