package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func checkDiff(a, b int) int {
	diff := a - b

	if diff > 0 {
		return 1
	} else if diff < 0 {
		return -1
	}
	return 0
}

func getUnSafeIndex(line []string) int {
	unsafe := 0
	report := line

	i := 0
	direction := 0 // 0 means start ,1 means increase , -1 means decrease
	prevVal := -10000
	for {
		val, _ := strconv.Atoi(report[i])

		if direction == 0 && i == 0 {
			// for cases when first item can be removed to make the report safe
			prevVal = val
		} else {
			diff := int(math.Abs(float64(val - prevVal)))
			if checkDiff(val, prevVal) == 1 && direction == -1 {
				unsafe = 1
			} else if checkDiff(val, prevVal) == -1 && direction == 1 {
				unsafe = 1
			} else if diff < 1 || diff > 3 {
				unsafe = 1
			} else {
				direction = checkDiff(val, prevVal)
			}
			prevVal = val
		}

		if unsafe == 1 {
			return i
		}

		i++
		if i == len(report) {
			break
		}
	}
	return -1
}

func isRepairable(line string) bool {

	delimiter := " "
	report := strings.Split(line, delimiter)

	// handle case when first item can be removed
	reportWithoutFirstLevel := report[1:]
	safetyCheckWithoutFirstLevel := getUnSafeIndex(reportWithoutFirstLevel)

	if safetyCheckWithoutFirstLevel == -1 {
		return true
	}

	firstUnsafeIndex := getUnSafeIndex(report)
	lineLength := len(report)

	report = append(report[:firstUnsafeIndex], report[firstUnsafeIndex+1:]...)

	if firstUnsafeIndex == lineLength-1 {
		return true
	}

	secondUnsafeIndex := getUnSafeIndex(report)

	if secondUnsafeIndex == -1 {
		return true
	}

	return false
}

func main() {
	filePath := "second.txt"
	lines, err := utils.ReadFileLineByLine(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	unsafe := 0
	countSafeReports := 0

	var unsafeLines []string

	for _, line := range lines {
		delimiter := " "
		report := strings.Split(line, delimiter)

		i := 0
		direction := 0 // 0 means start ,1 means increase , -1 means decrease
		prevVal := -10000
		for {
			val, _ := strconv.Atoi(report[i])

			if direction == 0 && i == 0 {
				prevVal = val
			} else {
				diff := int(math.Abs(float64(val - prevVal)))
				if checkDiff(val, prevVal) == 1 && direction == -1 {
					unsafe = 1
				} else if checkDiff(val, prevVal) == -1 && direction == 1 {
					unsafe = 1
				} else if diff < 1 || diff > 3 {
					unsafe = 1
				} else {
					direction = checkDiff(val, prevVal)
				}
				prevVal = val
			}

			i++
			if i == len(report) {
				break
			}
		}

		if unsafe != 1 {
			countSafeReports += 1
		} else {
			unsafeLines = append(unsafeLines, line)
			unsafe = 0
		}

	}

	for _, line := range unsafeLines {
		if isRepairable(line) == true {
			countSafeReports += 1
		}
	}

	fmt.Println(countSafeReports)

}
