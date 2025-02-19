package main

import "fmt"

func main() {
	board := [][]byte{	{5, 3, '.', '.', 7, '.', '.', '.', '.'},
						{6, '.', '.', 1, 9, 5, '.', '.', '.'},
						{'.', 9, 8, '.', '.', '.', '.', 6, '.'},
						{8, '.', '.', '.', 6, '.', '.', '.', 3},
						{4, '.', '.', 8, '.', 3, '.', '.', 1},
						{7, '.', '.', '.', 2, '.', '.', '.', 6},
						{'.', 6, '.', '.', '.', '.', 2, 8, '.'},
						{'.', '.', '.', 4, 1, 9, '.', '.', 5},
						{'.', '.', '.', '.', 8, '.', '.', 7, 9}}

	nullBoard := [][]byte{{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
						{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
						{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
						{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
						{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
						{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
						{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
						{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
						{'.', '.', '.', '.', '.', '.', '.', '.', '.'}}

	fmt.Println(isValidSudoku(board))
	fmt.Println(isValidSudoku(nullBoard))
}

func isValidSudoku(board [][]byte) bool {
	smallSquare := [9]map[byte]int{}
	lineVert := [9]map[byte]int{}
	lineHor := [9]map[byte]int{}
	var status bool

	for i := 0; i < 9; i++ {
		smallSquare[i] = make(map[byte]int)
		lineVert[i] = make(map[byte]int)
		lineHor[i] = make(map[byte]int)
	}

	for i, line := range board {
		for k, value := range line {
			if value == 46{
				continue
			}
			smallSquare[k/3+(i/3)*3], status = addElementToMap(value, smallSquare[k/3+(i/3)*3])
			if !status {
				return status
			}	
			lineVert[k], status = addElementToMap(value, lineVert[k])

			if !status {
				return status
			}

			lineHor[i], status = addElementToMap(value, lineHor[i])
			if !status {
				return status
			}
		}
	}

	return true
}

func addElementToMap(element byte, m map[byte]int) (map[byte]int, bool) {
	m[element]++
	if m[element] > 1 {
		return nil, false
	}
	return m, true
}