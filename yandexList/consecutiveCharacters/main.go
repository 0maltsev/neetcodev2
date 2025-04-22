package main

import()

func main() {}

func maxPower(s string) int {
    i, j:= 0, 0
    maxPower, currPower := 0, 0

    for i < len(s) && j < len(s) {
        if s[i] == s[j] {
            currPower++
            j++
        }else{
            if currPower > maxPower {
                maxPower = currPower
            }
            currPower = 0
            i = j
        }

    }
    if currPower > maxPower {
        maxPower = currPower
    }

    return maxPower
}
