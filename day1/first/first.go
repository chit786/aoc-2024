package main

import (
	"day1/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func calculateFirst(leftList, rightList []int) (int, error) {
	totalDifference := 0

	// loop over on sorted lists and pop first element while calculating total diff
	for i := 0; i < len(leftList); i++ {
		diff := int(math.Abs(float64(leftList[i] - rightList[i])))
		totalDifference += diff
	}
	return totalDifference, nil

}

func main() {
	filePath := "first.txt"
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

	// sorted list in ascending order
	sort.Ints(leftList)
	sort.Ints(rightList)

	fmt.Println(calculateFirst(leftList, rightList))

}
