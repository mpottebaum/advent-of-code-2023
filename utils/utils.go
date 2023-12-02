package utils

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFileToString(name string) string {
	str, err := os.ReadFile(name)

	if err != nil {
		fmt.Print(err)
	}

	return string(str)
}

func ParseInt(str string) (i int, err error) {
	parsedInt, err := strconv.ParseInt(str, 10, 64)
	i = int(parsedInt)
	return
}
