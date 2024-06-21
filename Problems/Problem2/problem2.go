package problem2

import (
	"strconv"
	"strings"
)

type Problem2 struct{}

func (p *Problem2) Solve(data string, part int) int {
	total := 0
	for _, line := range strings.Split(data, "\n") {
		if part == 1 {
			total += processLinePart1(&line)
		} else {
			total += processLinePart2(&line)
		}
	}
	return total
}

func processLinePart1(line *string) int {
	min, max := 0, 0

	for _, element := range strings.Fields(*line) {
		intVal, _ := strconv.Atoi(element)

		if min == 0 && max == 0 {
			min, max = intVal, intVal
		} else if intVal < min {
			min = intVal
		} else if intVal > max {
			max = intVal
		}
	}

	return max - min
}

func processLinePart2(line *string) int {
	values := mapToInts(strings.Fields(*line))
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[j]%values[i] == 0 {
				return values[j] / values[i]
			} else if values[i]%values[j] == 0 {
				return values[i] / values[j]
			}
		}
	}
	return 0
}

func mapToInts(stringVals []string) []int {
	intVals := make([]int, len(stringVals))
	for index, element := range stringVals {
		intVals[index], _ = strconv.Atoi(element)
	}
	return intVals
}
