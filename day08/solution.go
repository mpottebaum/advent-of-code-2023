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
	startEls := []string{}

	elMap := map[string]Instruct{}
	for i := 2; i < len(rows); i++ {
		row := rows[i]
		elAndInstr := strings.Split(row, " = ")
		el := elAndInstr[0]
		instrStr := elAndInstr[1]
		left := string([]rune(instrStr)[1:4])
		right := string([]rune(instrStr)[6:9])
		elMap[el] = Instruct{Left: left, Right: right}
		ender := rune(el[2])
		if ender == 'A' {
			startEls = append(startEls, el)
		}
	}

	steps := strings.Split(stepsStr, "")
	loopNums := []int{}
	for i := 0; i < len(startEls); i++ {
		startEl := startEls[i]
		el := startEl
		numSteps := 0
		steppedEls := map[string]bool{}
		var startLoopEl string
	stepper:
		for {
			for iStep := 0; iStep < len(steps); iStep++ {
				step := steps[iStep]
				instruct := elMap[el]
				if step == "L" {
					el = instruct.Left
				} else {
					el = instruct.Right
				}

				if el == startLoopEl {
					numSteps++
					break stepper
				}
				if len(startLoopEl) > 0 {
					numSteps++
				} else {
					_, steppedUpon := steppedEls[el]
					ender := rune(el[2])
					if steppedUpon && ender == 'Z' {
						startLoopEl = el
					} else {
						steppedEls[el] = true
					}
				}
			}
		}
		loopNums = append(loopNums, numSteps)
	}
	numSteps := utils.LCM(loopNums[0], loopNums[1], loopNums[2:]...)
	fmt.Println("Number or required steps: ", numSteps)
}
