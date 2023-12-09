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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
