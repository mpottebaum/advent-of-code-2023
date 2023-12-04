package day04

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func Solve(inputFile string) {
	fileName := "input"
	if len(inputFile) > 0 {
		fileName = inputFile
	}
	input := utils.ReadFileToString("day04/" + fileName + ".txt")

	scratchCards := strings.Split(input, "\n")
	totalPoints := 0

	// for each scratch card
	for cI := 0; cI < len(scratchCards); cI++ {
		scratchCard := scratchCards[cI]
		// isolate winning numbers and myNumbers
		cardTitleAndNums := strings.Split(scratchCard, ": ")
		allNums := cardTitleAndNums[1]
		winningAndMyNums := strings.Split(allNums, " | ")
		winningNums := strings.Split(winningAndMyNums[0], " ")
		myNums := strings.Split(winningAndMyNums[1], " ")
		// create lookup map of winning numbers
		var winningMap = map[string]bool{}
		for wI := 0; wI < len(winningNums); wI++ {
			winningNum := winningNums[wI]
			// account for extra space from single digits
			if len(winningNum) > 0 {
				winningMap[winningNum] = true
			}
		}
		// check each myNumber against lookup map
		points := 0
		for mI := 0; mI < len(myNums); mI++ {
			myNum := myNums[mI]
			_, ok := winningMap[myNum]
			if ok {
				// for each match, double point value (init points at 0)
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		totalPoints += points
	}

	fmt.Println("Total scratchcard points: ", totalPoints)
}
