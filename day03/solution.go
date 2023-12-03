package day03

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strings"
)

func CheckOtherRow(row string, prevI int, nextI int, regex regexp.Regexp) bool {
	rowChars := strings.Split(row, "")
	var iInit int
	if isPrevIValid := prevI >= 0; isPrevIValid {
		iInit = prevI
	} else {
		iInit = 0
	}
	var iEnd int
	if isNextIValid := (nextI + 1) < len(row); isNextIValid {
		iEnd = nextI + 1
	} else {
		iEnd = len(row)
	}
	for i := iInit; i < iEnd; i++ {
		rowChar := rowChars[i]
		if rowAdj := regex.MatchString(rowChar); rowAdj {
			return true
		}
	}
	return false
}

func Solve(inputFile string) {
	fileName := "input"
	if len(inputFile) > 0 {
		fileName = inputFile
	}
	input := utils.ReadFileToString("day03/" + fileName + ".txt")

	rows := strings.Split(input, "\n")
	digitFinder, digRegexErr := regexp.Compile("\\d")
	symbolFinder, symRegexErr := regexp.Compile("[^a-zA-Z0-9_\\.]")
	if digRegexErr != nil || symRegexErr != nil {
		fmt.Println("digit regex err ", digRegexErr)
		fmt.Println("symbol regex err ", symRegexErr)
	}
	sum := 0
	var prevRow string
	var nextRow string
	// for every row
	for rI := 0; rI < len(rows); rI++ {
		row := rows[rI]
		// find each number in the row
		rowChars := strings.Split(row, "")
		startI := 0
		endI := 1
		for startI < len(rowChars) {
			char := rowChars[startI]
			if digitFinder.MatchString(char) {
				numberChars := []string{char}
				// get rest of number
				for endI < len(rowChars) {
					nextChar := rowChars[endI]
					if digitFinder.MatchString(nextChar) {
						numberChars = append(numberChars, nextChar)
						endI++
					} else {
						break
					}
				}

				joinedNumber := strings.Join(numberChars, "")
				if numberInt, parseErr := utils.ParseInt(joinedNumber); parseErr == nil {
					// search for adjacent symbols
					isPartNumber := false
					prevI := startI - 1
					nextI := endI
					// a. search same row for adjacent symbols
					if prevI >= 0 {
						prevChar := rowChars[prevI]
						if prevAdj := symbolFinder.MatchString(prevChar); prevAdj {
							isPartNumber = true
						}
					}
					if isPartNumber == false && nextI < len(rowChars) {
						nextChar := rowChars[nextI]
						if nextAdj := symbolFinder.MatchString(nextChar); nextAdj {
							isPartNumber = true
						}
					}
					// b. search previous row for adjacent symbols
					if isPartNumber == false && rI-1 > 0 {
						//   	include diagonals
						prevRow = rows[rI-1]
						isPartNumber = CheckOtherRow(prevRow, prevI, nextI, *symbolFinder)
					}
					// c. search next row for adjacent symbols
					if isPartNumber == false && rI+1 < len(rows) {
						//   	include diagonals
						nextRow = rows[rI+1]
						isPartNumber = CheckOtherRow(nextRow, prevI, nextI, *symbolFinder)
					}
					// if any adj symbol is found, add number to sum
					if isPartNumber {
						sum += numberInt
					}
				}
				// set start index to next char after number search break
				startI = endI + 1
				endI = startI + 1
			} else {
				// no digit, increment indexes
				startI++
				endI = startI + 1
			}

		}
	}
	fmt.Println("Sum of all part numbers: ", sum)
}
