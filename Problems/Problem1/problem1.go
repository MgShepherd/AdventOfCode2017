package problem1

import (
	"strings"
)

type Problem1 struct{}

func (p *Problem1) Solve(data string, part int) int {
	matchingSum := 0
	data = strings.TrimSpace(data)
	for i := 0; i < len(data); i++ {
		compareIndex := getCompareIndex(i, part, len(data))
		if data[i] == data[compareIndex] {
			matchingSum += int(data[i] - '0')
		}
	}
	return matchingSum
}

func getCompareIndex(currentIndex, part, dataLength int) int {
	if part == 1 {
		return (currentIndex + 1) % dataLength
	}
	return (currentIndex + dataLength/2) % dataLength
}
