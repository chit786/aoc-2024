package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func countOccurences(list []int, target int) int {
	count := 0
	for _, num := range list {
		if num == target {
			count++
		}
	}
	return count
}

func calculateSecond(leftList, rightList []int) (int, error) {
	numberAndOccurenceMap := make(map[int]int)
	totalSimilarityScore := 0

	// loop over list
	for _, num := range leftList {
		value, exists := numberAndOccurenceMap[num]
		if exists {
			// just add into the similarity score
			totalSimilarityScore += (num * value)
		} else {
			// find occurence and store in the map at same time calculate the similarity score
			numOfOccurences := countOccurences(rightList, num)
			numberAndOccurenceMap[num] = numOfOccurences
			totalSimilarityScore += (num * numOfOccurences)
		}
	}

	fmt.Println(numberAndOccurenceMap)

	return totalSimilarityScore, nil

}

func main() {
	filePath := "second.txt"
	lines, err := utils.ReadFileLineByLine(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	var leftList []int
	var rightList []int

	for _, line := range lines {
		delimiter := "   "
		firstInput := strings.Split(line, delimiter)[0]
		secondInput := strings.Split(line, delimiter)[1]

		firstNum, err := strconv.Atoi(firstInput)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		secondNum, err := strconv.Atoi(secondInput)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		leftList = append(leftList, firstNum)
		rightList = append(rightList, secondNum)

	}

	fmt.Println(calculateSecond(leftList, rightList))

}
