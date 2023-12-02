package day02

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

var ElfBag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Solve() {
	input := utils.ReadFileToString("day02/input.txt")
	games := strings.Split(input, "\n")

	sum := 0
	for iG := 0; iG < len(games); iG++ {
		game := games[iG]
		idAndHandfuls := strings.Split(game, ": ")
		handfulsStr := idAndHandfuls[1]
		handfuls := strings.Split(handfulsStr, "; ")

		valid := true
	handfulCheck:
		for iH := 0; iH < len(handfuls); iH++ {
			handful := handfuls[iH]
			cubeCountStrs := strings.Split(handful, ", ")
			for iC := 0; iC < len(cubeCountStrs); iC++ {
				cubeCountStr := cubeCountStrs[iC]
				countAndColor := strings.Split(cubeCountStr, " ")
				count, countParseErr := strconv.ParseInt(countAndColor[0], 10, 64)
				color := countAndColor[1]
				colorTotal, colorOk := ElfBag[color]
				if countParseErr == nil && colorOk {
					if int(count) > colorTotal {
						valid = false
						break handfulCheck
					}
				}
			}
		}
		if valid {
			gameAndId := strings.Split(idAndHandfuls[0], " ")
			id, idParseErr := strconv.ParseInt(gameAndId[1], 10, 64)
			if idParseErr == nil {
				sum += int(id)
			}
		}
	}
	fmt.Println("Sum of all valid game IDs: ", sum)
}
