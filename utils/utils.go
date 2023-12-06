package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func ReadFileToString(name string) (str string) {
	b, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("ReadFileToString error: ", err)
	}
	str = string(b)
	return
}

func ParseInt(str string) (i int, err error) {
	parsedInt, err := strconv.ParseInt(str, 10, 64)
	i = int(parsedInt)
	return
}

func AutoQuadraticForThePeople(a, b, c float64) (float64, float64) {
	discriminant := (b * b) - (4 * a * c)
	rootA := (-b + math.Sqrt(discriminant)) / (2 * a)
	rootB := (-b - math.Sqrt(discriminant)) / (2 * a)
	return rootA, rootB
}
