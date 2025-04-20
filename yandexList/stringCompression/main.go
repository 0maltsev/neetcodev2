package main 

import (
    "strconv"
)

func main() {
	
}

func compress(chars []byte) int {
    readIndex, writeIndex := 0, 0
    for readIndex < len(chars) {
        currentChar := chars[readIndex]
        currentCount := 0

        // Count the number of consecutive repeating characters
        for readIndex < len(chars) && currentChar == chars[readIndex] {
            readIndex++
            currentCount++
        }

        // Write the compression result to chars
        chars[writeIndex] = currentChar
        writeIndex++

        if currentCount > 1 {
            count := strconv.Itoa(currentCount)
            for _, i := range(count) {
                chars[writeIndex] = byte(i)
                writeIndex++
            }
        }
    }
    return writeIndex
}