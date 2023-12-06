package day06

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
	input := utils.ReadFileToString("day06/" + fileName + ".txt")
	rows := strings.Split(input, "\n")
	// create slice of times, slice of distances from input
	timesStrs := strings.Fields(rows[0])
	distancesStrs := strings.Fields(rows[1])
	races := make([][2]int, 0, len(timesStrs)-1)
	for i := 0; i < len(timesStrs); i++ {
		timeStr := timesStrs[i]
		distanceStr := distancesStrs[i]
		timeInt, timeErr := utils.ParseInt(timeStr)
		distanceInt, distanceErr := utils.ParseInt(distanceStr)
		if timeErr == nil && distanceErr == nil {
			raceTuple := [2]int{
				timeInt,
				distanceInt,
			}
			races = append(races, raceTuple)
		}
	}
	multipliedWins := 1
	// for each race
	for i := 0; i < len(races); i++ {
		// test the different lengths of button holding
		race := races[i]
		time := race[0]
		distanceRecord := race[1]
		raceWins := 0
		for holdButtonMs := 1; holdButtonMs <= time; holdButtonMs++ {
			speed := holdButtonMs
			remainingMs := time - holdButtonMs
			distance := speed * remainingMs
			if distance > distanceRecord {
				raceWins++
			}
		}
		// multiply by and reassign to multipliedWins
		multipliedWins *= raceWins
	}
	fmt.Println("Number of ways to win each race multiplied by each other: ", multipliedWins)
}
