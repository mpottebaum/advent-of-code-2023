package main

import (
	day01 "aoc/day01"
	day02 "aoc/day02"
	"os"
)

func main() {
	if meArgs := os.Args; len(meArgs) == 2 {
		switch day := meArgs[1]; day {
		case "1":
			day01.Solve()
		case "2":
			day02.Solve()
		}
	}
}
