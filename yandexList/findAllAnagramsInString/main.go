package main

import()

func main() {

}

func findAnagrams(s string, p string) []int {
    var res []int
    
    if len(s) < len(p) {
        return res
    }
    
    charFrequency := make(map[byte]int)
    
    for i := 0; i < len(p); i++ {
        char := p[i]
        charFrequency[char]++
    }
    
    leftWindow, numberOfMatchChar := 0, 0
    
    for rightWindow := 0; rightWindow < len(s); rightWindow++ {
        for rightWindow - leftWindow + 1 > len(p) {
            leftChar := s[leftWindow]
            
            if frequency, ok := charFrequency[leftChar]; ok {
                if frequency == 0 {
                    numberOfMatchChar--
                }
                
                charFrequency[leftChar]++
            }
            
            leftWindow++
        }
        
        rightChar := s[rightWindow]
        
        if _, ok := charFrequency[rightChar]; ok {
            charFrequency[rightChar]--
            
            if charFrequency[rightChar] == 0 {
                numberOfMatchChar++
            }
        }
        
        if numberOfMatchChar == len(charFrequency) {
            res = append(res, leftWindow)
        }
    }
    
    return res
}
