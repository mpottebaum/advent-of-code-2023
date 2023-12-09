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
		sequences = append(sequences, values)
		// find diffs
		for {
			diffs := []int{}
			isBreaker := true
			for i := 0; i < len(values)-1; i++ {
				currVal := values[i]
				nextVal := values[i+1]
				diff := nextVal - currVal
				if diff != 0 && isBreaker {
					isBreaker = false
				}
				diffs = append(diffs, diff)
			}
			if isBreaker {
				break
			}

			sequences = append(sequences, diffs)
			values = diffs
		}

		// extrapolate row
		diff := 0
		for i := len(sequences) - 1; i >= 0; i-- {
			sequence := sequences[i]
			lastIdx := len(sequence) - 1
			lastVal := sequence[lastIdx]
			diff = lastVal + diff
		}
		// add extrapolated value to sum
		sum += diff
	}
	fmt.Println("Sum of extrapolated values: ", sum)
}
