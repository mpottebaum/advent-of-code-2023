package day08

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type Instruct struct {
	Left  string
	Right string
}

func Solve(inputFile string) {
	fileName := "input"
	if len(inputFile) > 0 {
		fileName = inputFile
	}
	input := utils.ReadFileToString("day08/" + fileName + ".txt")
	rows := strings.Split(input, "\n")
	stepsStr := rows[0]

	// collect all els + instrs in map
	elMap := map[string]Instruct{}
	for i := 2; i < len(rows); i++ {
		row := rows[i]
		elAndInstr := strings.Split(row, " = ")
		el := elAndInstr[0]
		instrStr := elAndInstr[1]
		left := string([]rune(instrStr)[1:4])
		right := string([]rune(instrStr)[6:9])
		elMap[el] = Instruct{Left: left, Right: right}
	}
	// walk the steps until we get to ZZZ
	steps := strings.Split(stepsStr, "")
	currEl := "AAA"
	numSteps := 0
	for currEl != "ZZZ" {
		for i := 0; i < len(steps); i++ {
			numSteps++
			instruct := elMap[currEl]
			if step := steps[i]; step == "L" {
				currEl = instruct.Left
			} else {
				currEl = instruct.Right
			}
			if currEl == "ZZZ" {
				break
			}
		}
	}
	fmt.Println("Number or required steps: ", numSteps)
}
