package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadProblemFile(problemNumber int) string {
	pwd, _ := os.Getwd()
	fileName := fmt.Sprintf("problem%d.txt", problemNumber)
	location := filepath.Join(pwd, "data", fileName)
	data, err := os.ReadFile(location)

	if err != nil {
		errorMessage := fmt.Sprintf("Unable to read file for problem %d, recieved error: %q", problemNumber, err)
		panic(errorMessage)
	}
	return string(data)
}
