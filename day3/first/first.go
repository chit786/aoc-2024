package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	filePath := "first.txt"
	lines, err := utils.ReadFileLineByLine(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	i := 0
	sumNumbers := 0

	fmt.Println(len(lines))

	var characters []string

	for _, line := range lines {
		characters = append(characters, strings.Split(line, "")...)
	}

	for i = 0; i < len(characters); i++ {
		for {
			currPointer := i

			if characters[i] == ")" {
				// found the closing bracket then traverse back until mul( is found
				rightNumberString := ""
				leftNumberString := ""
				for currPointer--; currPointer > 0; currPointer-- {
					_, err := strconv.Atoi(characters[currPointer])
					if err != nil {
						break
					}
					rightNumberString = characters[currPointer] + rightNumberString
				}

				if rightNumberString == "" {
					break
				}

				if characters[currPointer] == "," {
					for currPointer--; currPointer > 0; currPointer-- {
						_, err := strconv.Atoi(characters[currPointer])
						if err != nil {
							break
						}
						leftNumberString = characters[currPointer] + leftNumberString
					}

					if leftNumberString == "" {
						break
					}
				} else {
					break
				}
				fmt.Println("rightNumberString", rightNumberString)
				fmt.Println("leftNumberString", leftNumberString)

				if currPointer-3 > 0 {
					if characters[currPointer-3]+characters[currPointer-2]+characters[currPointer-1]+characters[currPointer] == "mul(" {
						leftNumber, err := strconv.Atoi(leftNumberString)
						if err != nil {
							fmt.Printf("Error converting left number: %v\n", err)
							return
						}

						rightNumber, err := strconv.Atoi(rightNumberString)
						if err != nil {
							fmt.Printf("Error converting right number: %v\n", err)
							return
						}

						sumNumbers += leftNumber * rightNumber
					}
				}

			}

			i++
			if i == len(characters) {
				break
			}
		}

	}

	fmt.Println(sumNumbers)

}
