package day01

import (
	utils "aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

var DigitIntMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var SpelledDigits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func SearchStrForInts(str string) (int, int) {
	chars := strings.Split(str, "")
	foundDigits := []int{}
	for chIdx := 0; chIdx < len(chars); chIdx++ {
		char := chars[chIdx]
		if digitInt, err := strconv.ParseInt(char, 10, 64); err == nil {
			foundDigits = append(foundDigits, int(digitInt))
		} else {
		spelledCheck:
			for spIdx := 0; spIdx < len(SpelledDigits); spIdx++ {
				spelledDigit := SpelledDigits[spIdx]
				if endOfSlice := chIdx + len(spelledDigit); endOfSlice <= len(chars) {
					possibleSpelled := strings.Join(chars[chIdx:endOfSlice], "")
					if spelledDigitInt, isSpelled := DigitIntMap[possibleSpelled]; isSpelled {
						foundDigits = append(foundDigits, spelledDigitInt)
						break spelledCheck
					}
				}
			}
		}
	}
	firstDigit := foundDigits[0]
	lastDigit := foundDigits[len(foundDigits)-1]
	return firstDigit, lastDigit
}

func JoinInts(first, last int) int {
	firstStr := strconv.Itoa(first)
	lastStr := strconv.Itoa(last)
	var sb strings.Builder
	sb.WriteString(firstStr)
	sb.WriteString(lastStr)
	joinedInt, err := strconv.ParseInt(sb.String(), 10, 64)
	if err != nil {
		fmt.Println("JoinInts parse error:", err)
		return 0
	}
	return int(joinedInt)
}

func Solve() {
	input := utils.ReadFileToString("day01/input.txt")
	rows := strings.Split(input, "\n")
	var sum int
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		firstInt, lastInt := SearchStrForInts(row)
		joinedInt := JoinInts(firstInt, lastInt)
		sum += joinedInt
	}
	fmt.Println("Sum of all calibration values", sum)
}
