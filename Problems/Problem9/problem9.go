package problem9

import (
	"strings"
)

type Problem9 struct{}

func (p *Problem9) Solve(data string, part int) int {
	totalScore, garbageCount := 0, 0
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}
		lineScore, lineGarbage := getLineScoreAndGarbage(line)
		totalScore += lineScore
		garbageCount += lineGarbage
	}

	if part == 1 {
		return totalScore
	}
	return garbageCount
}

func getLineScoreAndGarbage(line string) (int, int) {
	nestLevel := 1
	currentIndex, lineScore, lineGarbage := 0, 0, 0
	inGarbage := false

	for currentIndex < len(line) {
		if line[currentIndex] == '!' {
			currentIndex += 1
		} else if line[currentIndex] == '>' {
			inGarbage = false
		} else if inGarbage {
			lineGarbage += 1
		} else {
			switch line[currentIndex] {
			case '{':
				lineScore += nestLevel
				nestLevel += 1
			case '}':
				nestLevel -= 1
			case '<':
				inGarbage = true
			}
		}
		currentIndex += 1
	}
	return lineScore, lineGarbage
}
