package main

import(
	"fmt"
)

func main() {
	digs := []int{1, 2, 3}
	fmt.Println(plusOne(digs))
}

func plusOne(digits []int) []int {
    one := 1
    i := 0
    digits = reverse(digits)

    for one != 0 {
        if i < len(digits) {
            if digits[i] == 9 {
                digits[i] = 0
            } else {
                digits[i] += 1
                one = 0
            }
        } else {
            digits = append(digits, one)
            one = 0
        }
        i++
    }
    return reverse(digits)
}

func reverse(digits []int) []int {
    for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
        digits[i], digits[j] = digits[j], digits[i]
    }
    return digits
}