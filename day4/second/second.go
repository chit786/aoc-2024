package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func main() {
	filePath := "second.txt"
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

	countMAS := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			if matrix[i][j] == "A" {

				// backslash

				if i-1 >= 0 && j-1 >= 0 && i+1 < len(matrix) && j+1 < len(matrix[i]) {

					str1 := matrix[i-1][j-1] + "A" + matrix[i+1][j+1]

					// rightslash

					str2 := matrix[i-1][j+1] + "A" + matrix[i+1][j-1]

					if (str1 == "SAM" || str1 == "MAS") && (str2 == "SAM" || str2 == "MAS") {
						countMAS += 1
					}

				}
			}

		}
	}

	fmt.Println(countMAS)

}
