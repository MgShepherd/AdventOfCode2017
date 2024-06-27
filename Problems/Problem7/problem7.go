package problem7

import (
	"fmt"
	"strconv"
	"strings"
)

type Problem7 struct{}

func (p *Problem7) Solve(data string, part int) string {
	if part == 1 {
		return solvePart1(data)
	}
	return solvePart2(data)
}

func solvePart1(data string) string {
	parentNodes := make(map[string]string)
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}
		parent, _, children := convertLineToTokens(&line)
		for _, child := range children {
			parentNodes[child] = parent
		}
		if _, exists := parentNodes[parent]; !exists {
			parentNodes[parent] = ""
		}
	}
	return findBottomTower(parentNodes)
}

func solvePart2(data string) string {
	bottomTowerName := solvePart1(data)
	childNodes := make(map[string][]string)
	weights := make(map[string]int)
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}
		parent, parentWeight, children := convertLineToTokens(&line)
		childNodes[parent] = children
		weights[parent] = parentWeight
	}
	return getRequiredWeightChange(childNodes, weights, bottomTowerName)
}

func convertLineToTokens(line *string) (string, int, []string) {
	*line = strings.TrimSpace(*line)
	tokens := strings.Fields(*line)
	var children []string
	if len(tokens) > 3 {
		for i := 3; i < len(tokens); i++ {
			tokens[i] = strings.Trim(tokens[i], ",")
			children = append(children, tokens[i])
		}
	}
	weight, _ := strconv.Atoi(strings.Trim(tokens[1], "()\\n"))
	return tokens[0], weight, children
}

func findBottomTower(parentNodes map[string]string) string {
	for key, value := range parentNodes {
		if value == "" {
			return key
		}
	}
	return "Unknown"
}

func getRequiredWeightChange(childNodes map[string][]string, weights map[string]int, bottomTower string) string {
	_, weightChange := getWeights(bottomTower, childNodes, weights)
	return fmt.Sprintf("%d", weightChange)
}

func getWeights(currentNode string, childNodes map[string][]string, weights map[string]int) (int, int) {
	if len(childNodes[currentNode]) == 0 {
		return weights[currentNode], 0
	}
	var childWeights []int
	for _, node := range childNodes[currentNode] {
		weight, change := getWeights(node, childNodes, weights)
		if change != 0 {
			return 0, change
		}
		childWeights = append(childWeights, weight)
	}

	differentWeightIndex := getDifferentWeightedTowerIndex(childWeights)
	if differentWeightIndex != -1 {
		if childWeights[differentWeightIndex] < childWeights[(differentWeightIndex+1)%len(childWeights)] {
			return 0, weights[childNodes[currentNode][differentWeightIndex]] +
				(childWeights[(differentWeightIndex+1)%len(childWeights)] - childWeights[differentWeightIndex])
		}
		return 0, weights[childNodes[currentNode][differentWeightIndex]] -
			(childWeights[differentWeightIndex] - childWeights[(differentWeightIndex+1)%len(childWeights)])
	}

	weightSum := 0
	for i := 0; i < len(childWeights); i++ {
		weightSum += childWeights[i]
	}
	return weights[currentNode] + weightSum, 0
}

func getDifferentWeightedTowerIndex(childWeights []int) int {
	weightIndexes := make(map[int][]int)
	for index, weight := range childWeights {
		weightIndexes[weight] = append(weightIndexes[weight], index)
	}

	if len(weightIndexes) != 1 {
		for _, val := range weightIndexes {
			if len(val) == 1 {
				return val[0]
			}
		}
	}
	return -1
}
