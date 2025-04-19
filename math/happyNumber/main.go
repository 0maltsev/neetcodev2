package main

import(
	"fmt"
)

func main() {
	num := 1
	fmt.Println(isHappy(num))
}

func isHappy(n int) bool {
    slow, fast := n, sumOfSquares(n)

    for slow != fast {
        fast = sumOfSquares(fast)
        fast = sumOfSquares(fast)
        slow = sumOfSquares(slow)
    }

    return fast == 1
}

func sumOfSquares(n int) int {
    output := 0
    for n > 0 {
        digit := n % 10
        output += digit * digit
        n = n / 10
    }
    return output
}