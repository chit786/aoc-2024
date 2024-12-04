package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func main() {
	filePath := "first.txt"
	lines, err := utils.ReadFileLineByLine(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	rows := len(lines)
	cols := len(strings.Split(lines[0], ""))

	matrix := make([][]string, rows)
	for i := range matrix {
		matrix[i] = make([]string, cols)
	}

	for i, line := range lines {
		for j, char := range strings.Split(line, "") {
			matrix[i][j] = char
		}
	}

	countXMAS := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			if matrix[i][j] == "X" {

				// check up
				if i-3 >= 0 {
					if matrix[i-1][j]+matrix[i-2][j]+matrix[i-3][j] == "MAS" {
						countXMAS += 1
					}
				}

				// check left
				if j-3 >= 0 {
					if matrix[i][j-1]+matrix[i][j-2]+matrix[i][j-3] == "MAS" {
						countXMAS += 1
					}
				}

				// check down
				if i+3 < len(matrix) {
					if matrix[i+1][j]+matrix[i+2][j]+matrix[i+3][j] == "MAS" {
						countXMAS += 1
					}
				}

				// check right
				if j+3 < len(matrix[i]) {
					if matrix[i][j+1]+matrix[i][j+2]+matrix[i][j+3] == "MAS" {
						countXMAS += 1
					}
				}

				// check left up

				if j-3 >= 0 && i-3 >= 0 {
					if matrix[i-1][j-1]+matrix[i-2][j-2]+matrix[i-3][j-3] == "MAS" {
						countXMAS += 1
					}
				}

				// check right up

				if j+3 < len(matrix[i]) && i-3 >= 0 {
					if matrix[i-1][j+1]+matrix[i-2][j+2]+matrix[i-3][j+3] == "MAS" {
						countXMAS += 1
					}
				}

				// check right down

				if j+3 < len(matrix[i]) && i+3 < len(matrix) {
					if matrix[i+1][j+1]+matrix[i+2][j+2]+matrix[i+3][j+3] == "MAS" {
						countXMAS += 1
					}
				}

				// check left down

				if j-3 >= 0 && i+3 < len(matrix) {
					if matrix[i+1][j-1]+matrix[i+2][j-2]+matrix[i+3][j-3] == "MAS" {
						countXMAS += 1
					}
				}

			}

		}
	}

	fmt.Println(countXMAS)

}
