package problem4

import (
	"sort"
	"strings"
)

type Problem4 struct{}

func (p *Problem4) Solve(data string, part int) int {
	validPassphrases := 0
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}
		if (part == 1 && isPhraseValidPart1(line)) || (part == 2 && isPhraseValidPart2(line)) {
			validPassphrases += 1
		}
	}
	return validPassphrases
}

func isPhraseValidPart1(phrase string) bool {
	elements := make(map[string]int)

	for _, element := range strings.Fields(phrase) {
		if elements[element] != 0 {
			return false
		}
		elements[element] = 1
	}
	return true
}

func isPhraseValidPart2(phrase string) bool {
	elements := make(map[string]int)

	for _, element := range strings.Fields(phrase) {
		sortedElement := convertToSortedString(element)
		if elements[sortedElement] != 0 {
			return false
		}
		elements[sortedElement] = 1
	}
	return true
}

func convertToSortedString(value string) string {
	runes := []rune(value)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
