package day06

import (
	"aoc/utils"
	"fmt"
	"math"
	"strings"
)

func Solve(inputFile string) {
	fileName := "input"
	if len(inputFile) > 0 {
		fileName = inputFile
	}
	input := utils.ReadFileToString("day06/" + fileName + ".txt")
	rows := strings.Split(input, "\n")

	timesStrs := strings.Fields(rows[0])
	timeStr := strings.Join(timesStrs[1:], "")
	distancesStrs := strings.Fields(rows[1])
	distanceStr := strings.Join(distancesStrs[1:], "")

	var wins int

	time, timeErr := utils.ParseInt(timeStr)
	distanceRecord, distanceErr := utils.ParseInt(distanceStr)

	if timeErr == nil && distanceErr == nil {
		a := float64(-1)
		b := float64(time)
		c := float64(-1 * distanceRecord)

		discriminant := (b * b) - (4 * a * c)
		rootAFloat := (-b + math.Sqrt(discriminant)) / (2 * a)
		rootBFloat := (-b - math.Sqrt(discriminant)) / (2 * a)
		minHold := math.Floor(rootAFloat) + 1
		maxHold := math.Ceil(rootBFloat) - 1

		wins = int(maxHold - minHold + 1)
	}
	fmt.Println("Number of ways to win: ", wins)
}
