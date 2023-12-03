package day03

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func SearchForDigits(i int, rowChars []string, isInBounds bool, digitFinder regexp.Regexp) (string, bool) {
	var char string
	if isInBounds {
		char = rowChars[i]
		if char := rowChars[i]; digitFinder.MatchString(char) {
			return char, true
		}
	}
	return char, false
}

func GetDigit(rowChars []string, startI int, digitFinder regexp.Regexp) (int, error) {
	startChar := rowChars[startI]
	leftChars := []string{}
	rightChars := []string{}

	isLeftActive := true
	isRightActive := true
	left := startI - 1
	right := startI + 1
	for isLeftActive || isRightActive {
		leftChar, isLeftDigit := SearchForDigits(left, rowChars, left >= 0, digitFinder)
		if isLeftDigit {
			leftChars = append(leftChars, leftChar)
			left--
		}
		isLeftActive = isLeftDigit

		rightChar, isRightDigit := SearchForDigits(right, rowChars, right < len(rowChars), digitFinder)
		if isRightDigit {
			rightChars = append(rightChars, rightChar)
			right++
		}
		isRightActive = isRightDigit
	}
	slices.Reverse(leftChars)
	numberChars := append(leftChars, startChar)
	numberChars = append(numberChars, rightChars...)
	joinedNumber := strings.Join(numberChars, "")
	numberInt, parseErr := utils.ParseInt(joinedNumber)
	return numberInt, parseErr
}

func GetAdjacentDigitFromRow(row string, startI int, digitFinder regexp.Regexp) []int {
	rowChars := strings.Split(row, "")
	adjDigits := []int{}

	for i := startI; ; {
		if i >= 0 && i < len(rowChars) {
			char := rowChars[i]
			if digitFinder.MatchString(char) {
				if digit, err := GetDigit(rowChars, i, digitFinder); err == nil {
					adjDigits = append(adjDigits, digit)
					if i == startI {
						return adjDigits
					}
				}
			}
		}
		if i == startI {
			i = startI + 1
		} else if i == startI+1 {
			i = startI - 1
		} else {
			break
		}
	}
	return adjDigits
}

func Solve(inputFile string) {
	fileName := "input"
	if len(inputFile) > 0 {
		fileName = inputFile
	}
	input := utils.ReadFileToString("day03/" + fileName + ".txt")

	rows := strings.Split(input, "\n")
	digitFinder, digRegexErr := regexp.Compile("\\d")
	if digRegexErr != nil {
		fmt.Println("digit regex err ", digRegexErr)
	}
	sum := 0
	var prevRow string
	var nextRow string
	var row string
	var rowChars []string
	var char string
	var startI int
	var prevI int
	var nextI int
	var adjacentNumbers []int
	// for every row
	for rI := 0; rI < len(rows); rI++ {
		row = rows[rI]
		// find each number in the row
		rowChars = strings.Split(row, "")
		startI = 0
		for ; startI < len(rowChars); startI++ {
			char = rowChars[startI]
			if char == "*" {
				// search for adjacent numbers
				prevI = startI - 1
				nextI = startI + 1
				adjacentNumbers = []int{}
				// a. search same row for adjacent numbers
				if prevI >= 0 {
					if prevInt, err := GetDigit(rowChars, prevI, *digitFinder); err == nil {
						adjacentNumbers = append(adjacentNumbers, prevInt)
					}
				}
				if nextI < len(rowChars) {
					if nextInt, err := GetDigit(rowChars, nextI, *digitFinder); err == nil {
						adjacentNumbers = append(adjacentNumbers, nextInt)
					}
				}
				// b. search previous row for adjacent numbers
				if rI-1 >= 0 {
					prevRow = rows[rI-1]
					moreAdjNums := GetAdjacentDigitFromRow(prevRow, startI, *digitFinder)
					if len(moreAdjNums) > 0 {
						adjacentNumbers = append(adjacentNumbers, moreAdjNums...)
					}
				}
				// c. search next row for adjacent numbers
				if rI+1 < len(rows) {
					nextRow = rows[rI+1]
					moreAdjNums := GetAdjacentDigitFromRow(nextRow, startI, *digitFinder)
					if len(moreAdjNums) > 0 {
						adjacentNumbers = append(adjacentNumbers, moreAdjNums...)
					}
				}
				// if there are two adjacent numbers, multiply them and add to sum
				if len(adjacentNumbers) == 2 {
					partNumOne := adjacentNumbers[0]
					partNumTwo := adjacentNumbers[1]
					gearRatio := partNumOne * partNumTwo
					sum += gearRatio
				}
			}
		}
	}
	fmt.Println("Sum of all part numbers: ", sum)
}
