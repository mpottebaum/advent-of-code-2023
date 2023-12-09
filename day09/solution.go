package day08

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
	input := utils.ReadFileToString("day09/" + fileName + ".txt")
	rows := strings.Split(input, "\n")

	// for each row
	sum := 0
	for _, row := range rows {
		valueStrs := strings.Split(row, " ")
		sequences := [][]int{}
		values := []int{}
		for _, valStr := range valueStrs {
			if value, err := utils.ParseInt(valStr); err == nil {
				values = append(values, value)
			}
		}
		// find diffs
		diffs := values
		for {
			sequences = append(sequences, diffs)
			values = diffs
			diffs = []int{}
			isZeroes := true
			for i := 0; i < len(values)-1; i++ {
				currVal := values[i]
				nextVal := values[i+1]
				diff := nextVal - currVal
				if diff != 0 && isZeroes {
					isZeroes = false
				}
				diffs = append(diffs, diff)
			}
			if isZeroes {
				break
			}
		}
		// extrapolate row
		diff := 0
		for i := len(sequences) - 1; i >= 0; i-- {
			sequence := sequences[i]
			firstVal := sequence[0]
			diff = firstVal - diff
		}
		// add extrapolated value to sum
		sum += diff
	}
	fmt.Println("Sum of extrapolated values: ", sum)
}
