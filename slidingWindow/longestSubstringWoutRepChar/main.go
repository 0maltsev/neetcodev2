package main

import ("fmt")

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
    charSet := make(map[byte]bool)
    left, res := 0, 0

    for right := 0; right < len(s); right++ {
        for charSet[s[right]] {
            delete(charSet, s[left])
            left++
        }
        charSet[s[right]] = true
        if right - left + 1 > res {
            res = right - left + 1
        }
    }
    return res
}