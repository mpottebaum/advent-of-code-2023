package day10

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type Position struct {
	Row       int
	Col       int
	Direction string
}

var DirectionConnections = map[string]map[rune]string{
	"up": {
		'|': "up",
		'F': "right",
		'7': "left",
	},
	"down": {
		'|': "down",
		'L': "right",
		'J': "left",
	},
	"left": {
		'-': "left",
		'F': "down",
		'L': "up",
	},
	"right": {
		'-': "right",
		'7': "down",
		'J': "up",
	},
}

func FindNextPosition(currPos Position, rows []string) (Position, bool) {
	var nextDir string
	nextRow := currPos.Row
	nextCol := currPos.Col
	switch currPos.Direction {
	case "up":
		nextRow = currPos.Row - 1
	case "down":
		nextRow = currPos.Row + 1
	case "left":
		nextCol = currPos.Col - 1
	case "right":
		nextCol = currPos.Col + 1
	}
	if nextRow >= 0 && nextRow < len(rows) && nextCol >= 0 {
		if row := rows[nextRow]; nextCol < len(row) {
			connections := DirectionConnections[currPos.Direction]
			rowRunes := []rune(row)
			positionRune := rowRunes[nextCol]
			nextDir = connections[positionRune]
		}
	}
	return Position{Row: nextRow, Col: nextCol, Direction: nextDir}, len(nextDir) > 0
}

func Solve(inputFile string) {
	fileName := "input"
	if len(inputFile) > 0 {
		fileName = inputFile
	}
	input := utils.ReadFileToString("day10/" + fileName + ".txt")
	rows := strings.Split(input, "\n")

	// find the start position => row and column
	var startPosition Position
	for iRow, row := range rows {
		rowRunes := []rune(row)
		for iCol, rowRune := range rowRunes {
			if rowRune == 'S' {
				startPosition = Position{Row: iRow, Col: iCol}
				break
			}
		}
	}
	// find the two paths
	possibles := [4]string{
		"up",
		"down",
		"left",
		"right",
	}
	paths := []Position{}
	var isValid bool
	position := Position{}
	for _, direction := range possibles {
		position.Row = startPosition.Row
		position.Col = startPosition.Col
		position.Direction = direction
		position, isValid = FindNextPosition(position, rows)
		if isValid {
			paths = append(paths, position)
		}
	}

	// follow each path until they meet
	steps := 1
	for (paths[0].Row == paths[1].Row && paths[0].Col == paths[1].Col) == false {
		paths[0], _ = FindNextPosition(paths[0], rows)
		paths[1], _ = FindNextPosition(paths[1], rows)
		steps++
	}
	fmt.Println("Steps to farthest point: ", steps)
}
