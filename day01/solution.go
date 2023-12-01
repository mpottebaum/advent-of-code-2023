package day01

import (
	utils "aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func SearchStrForInts(str string) (int, int) {
	chars := strings.Split(str, "")
	var firstDigit, lastDigit int
	strLen := len(chars)
	start := 0
	end := strLen - 1
	for firstDigit == 0 || lastDigit == 0 {

		startChar := chars[start]
		endChar := chars[end]

		if startInt, err := strconv.ParseInt(startChar, 10, 64); firstDigit == 0 && err == nil {
			firstDigit = int(startInt)
		}

		if endInt, err := strconv.ParseInt(endChar, 10, 64); lastDigit == 0 && err == nil {
			lastDigit = int(endInt)
		}
		if firstDigit == 0 {
			start++
		}
		if lastDigit == 0 {
			end--
		}
	}
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
