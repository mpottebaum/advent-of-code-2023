package day02

import (
	"aoc/utils"
	"fmt"
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

		maxElfBag := map[string]int{}
		for iH := 0; iH < len(handfuls); iH++ {
			handful := handfuls[iH]
			cubeCountStrs := strings.Split(handful, ", ")
			for iC := 0; iC < len(cubeCountStrs); iC++ {
				cubeCountStr := cubeCountStrs[iC]
				countAndColor := strings.Split(cubeCountStr, " ")
				if count, countParseErr := utils.ParseInt(countAndColor[0]); countParseErr == nil {
					color := countAndColor[1]
					colorTotal, colorOk := maxElfBag[color]
					if (colorOk && colorTotal < count) || !colorOk {
						maxElfBag[color] = count
					}
				}
			}
		}
		power := 1
		for _, maxCount := range maxElfBag {
			power *= maxCount
		}
		sum += power
	}
	fmt.Println("Sum of every set power: ", sum)
}
