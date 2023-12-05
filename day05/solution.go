package day05

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
	input := utils.ReadFileToString("day05/" + fileName + ".txt")

	// split input string into: seeds list + all maps data
	splitness := strings.Split(input, "\n")
	seedsStr := splitness[0]
	// compile slice of maps data
	// structure == [[slice of map range data [num, num, num]]]
	// [][][num, num, num]
	mapsDataSlice := [][][3]int{}
	currentMapIndex := -1
	for i := 1; i < len(splitness); i++ {
		mapStr := splitness[i]
		if len(mapStr) > 0 {
			// check if str starts parseable int or map name
			mapStrValues := strings.Split(mapStr, " ")
			firstValueStr := mapStrValues[0]

			if firstInt, parseErr := utils.ParseInt(firstValueStr); parseErr == nil && currentMapIndex >= 0 {
				// starts with int
				newRangeTuple := [3]int{firstInt}
				for intI := 1; intI < len(mapStrValues); intI++ {
					nextInt, _ := utils.ParseInt(mapStrValues[intI])
					newRangeTuple[intI] = nextInt
				}
				mapsDataSlice[currentMapIndex] = append(mapsDataSlice[currentMapIndex], newRangeTuple)
			} else if len(firstValueStr) > 0 {
				// starts with letter
				// initialize map strings slice
				currentMapIndex++
				newDataSlice := [][3]int{}
				mapsDataSlice = append(mapsDataSlice, newDataSlice)
			}
		}
	}
	// search through seeds to find lowest location num
	seedNumStrs := strings.Split(seedsStr, " ")
	lowestLocationNum := -1
	// start at index 1 b/c 0 is "seeds:" str
	for sI := 1; sI < len(seedNumStrs); sI++ {
		seedNumStr := seedNumStrs[sI]
		seedNum, numParseErr := utils.ParseInt(seedNumStr)
		if numParseErr == nil {
			num := seedNum
			for mapI := 0; mapI < len(mapsDataSlice); mapI++ {
				// for each map
				mapness := mapsDataSlice[mapI]
				// find all source ranges
				for mapDatumI := 0; mapDatumI < len(mapness); mapDatumI++ {
					mapDatum := mapness[mapDatumI]
					sourceRangeStart := mapDatum[1]
					rangeLength := mapDatum[2]
					// if num fall within a range, use map to find new num
					if num >= sourceRangeStart && num < sourceRangeStart+rangeLength {
						destinationRangeStart := mapDatum[0]
						distanceFromRangeStart := num - sourceRangeStart
						num = destinationRangeStart + distanceFromRangeStart
						break
					}
					// else num stays the same
				}
			}
			// if num is less than lowest so far, replace lowest so far
			if lowestLocationNum == -1 || (lowestLocationNum >= 0 && num < lowestLocationNum) {
				lowestLocationNum = num
			}
		}
	}
	fmt.Println("Lowest location number", lowestLocationNum)
}
