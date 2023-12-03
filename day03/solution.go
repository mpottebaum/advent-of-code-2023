package day03

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func GetDigit(rowChars []string, startI int, digitFinder regexp.Regexp) (int, error) {
	startChar := rowChars[startI]
	leftChars := []string{}
	rightChars := []string{}

	isLeftActive := true
	isRightActive := true
	left := startI - 1
	right := startI + 1
	for isLeftActive || isRightActive {
		if isLeftActive == true && left >= 0 {
			leftChar := rowChars[left]
			if isDigit := digitFinder.MatchString(leftChar); isDigit {
				leftChars = append(leftChars, leftChar)
			} else {
				isLeftActive = false
			}
			left--
		} else {
			isLeftActive = false
		}
		if isRightActive == true && right < len(rowChars) {
			rightChar := rowChars[right]
			if isDigit := digitFinder.MatchString(rightChar); isDigit {
				rightChars = append(rightChars, rightChar)
			} else {
				isRightActive = false
			}
			right++
		} else {
			isRightActive = false
		}
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
	rowChar := rowChars[startI]
	adjDigits := []int{}
	if isDigit := digitFinder.MatchString(rowChar); isDigit {
		if digit, err := GetDigit(rowChars, startI, digitFinder); err == nil {
			adjDigits = append(adjDigits, digit)
			return adjDigits
		}
	}
	if prevI := startI - 1; prevI >= 0 {
		if digit, err := GetDigit(rowChars, prevI, digitFinder); err == nil {
			adjDigits = append(adjDigits, digit)
		}
	}
	if nextI := startI + 1; nextI < len(rowChars) {
		if digit, err := GetDigit(rowChars, nextI, digitFinder); err == nil {
			adjDigits = append(adjDigits, digit)
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
	// for every row
	for rI := 0; rI < len(rows); rI++ {
		row := rows[rI]
		// find each number in the row
		rowChars := strings.Split(row, "")
		for startI := 0; startI < len(rowChars); startI++ {
			char := rowChars[startI]
			if char == "*" {
				// search for adjacent numbers
				prevI := startI - 1
				nextI := startI + 1
				adjacentNumbers := []int{}
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
