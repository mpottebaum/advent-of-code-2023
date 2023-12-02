package utils

import (
	"fmt"
	"os"
)

func ReadFileToString(name string) string {
	b, err := os.ReadFile(name)

	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}
