package problem6

import (
	"strconv"
	"strings"
)

type Problem6 struct{}

func (p *Problem6) Solve(data string, part int) int {
	banks := convertToIntSlice(&data)
	seenConfigurations := make(map[string]int)
	numSteps := 0

	for {
		key := convertToKeyString(banks)
		if seenConfigurations[key] != 0 {
			if part == 2 {
				return numSteps - seenConfigurations[key]
			}
			break
		}

		seenConfigurations[key] = numSteps
		maxIndex, maxVal := findMax(banks)
		distributeBlocks(banks, maxIndex, maxVal)
		numSteps += 1
	}

	return numSteps
}

func convertToIntSlice(data *string) []int {
	var intSlice []int
	for _, item := range strings.Fields(*data) {
		if len(item) == 0 {
			continue
		}

		intVal, _ := strconv.Atoi(item)
		intSlice = append(intSlice, intVal)
	}
	return intSlice
}

func convertToKeyString(data []int) string {
	output := ""
	for _, element := range data {
		output += strconv.Itoa(element)
	}
	return output
}

func findMax(banks []int) (int, int) {
	maxIndex, maxVal := 0, 0

	for index, value := range banks {
		if value > maxVal {
			maxIndex = index
			maxVal = value
		}
	}

	return maxIndex, maxVal
}

func distributeBlocks(banks []int, maxIndex, value int) {
	banks[maxIndex] = 0
	currentIndex := (maxIndex + 1) % len(banks)
	for value > 0 {
		banks[currentIndex] += 1
		currentIndex = (currentIndex + 1) % len(banks)
		value -= 1
	}
}
