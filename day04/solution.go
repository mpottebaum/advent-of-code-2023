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
	copiesMap := map[int]int{}

	var scratchCard, winningNum, myNum string
	var cardTitleAndNums, winningAndMyNums, winningNums, myNums []string
	var winningMap map[string]bool
	var matchingNums, cardNum, nextNum, numCopies int
	// for each scratch card
	for cI := 0; cI < len(scratchCards); cI++ {
		scratchCard = scratchCards[cI]
		// winning numbers and myNumbers
		cardTitleAndNums = strings.Split(scratchCard, ": ")
		winningAndMyNums = strings.Split(cardTitleAndNums[1], " | ")
		winningNums = strings.Split(winningAndMyNums[0], " ")
		myNums = strings.Split(winningAndMyNums[1], " ")
		// create lookup map of winning numbers
		winningMap = map[string]bool{}
		for wI := 0; wI < len(winningNums); wI++ {
			// account for extra space from single digits
			if winningNum = winningNums[wI]; len(winningNum) > 0 {
				winningMap[winningNum] = true
			}
		}
		// check each myNumber against lookup map
		matchingNums = 0
		for mI := 0; mI < len(myNums); mI++ {
			myNum = myNums[mI]
			if _, isMatch := winningMap[myNum]; isMatch {
				// count matches
				matchingNums += 1
			}
		}
		// tally card copies from win
		cardNum = cI + 1
		numCopies = copiesMap[cardNum]
		for nextNum = cardNum + 1; nextNum <= cardNum+matchingNums && nextNum <= len(scratchCards); nextNum++ {
			if _, exists := copiesMap[nextNum]; exists {
				copiesMap[nextNum] += 1 + numCopies
			} else {
				copiesMap[nextNum] = 1 + numCopies
			}
		}
	}
	totalCardCount := len(scratchCards)
	for _, copiesCount := range copiesMap {
		totalCardCount += copiesCount
	}
	fmt.Println("Total scratchcard count: ", totalCardCount)
}
