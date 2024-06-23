package problem5

import (
	"strconv"
	"strings"
)

type Problem5 struct{}

func (p *Problem5) Solve(data string, part int) int {
	elements := convertToIntSlice(&data)
	currentInstruction := 0
	numSteps := 0
	for currentInstruction < len(elements) {
		processInstruction(&currentInstruction, elements, part)
		numSteps += 1
	}
	return numSteps
}

func convertToIntSlice(data *string) []int {
	var output []int
	for _, element := range strings.Split(*data, "\n") {
		element = strings.TrimSpace(element)
		if len(element) == 0 {
			continue
		}

		intVal, err := strconv.Atoi(element)
		if err != nil {
			panic("Unable to parse string to int")
		}

		output = append(output, intVal)
	}
	return output
}

func processInstruction(currentInstruction *int, elements []int, part int) {
	incrementValue := elements[*currentInstruction]
	if part == 1 || (part == 2 && incrementValue < 3) {
		elements[*currentInstruction] += 1
	} else {
		elements[*currentInstruction] -= 1
	}
	*currentInstruction = *currentInstruction + incrementValue
}
