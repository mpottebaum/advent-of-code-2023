package utils

import (
	"fmt"
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
