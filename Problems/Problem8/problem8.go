package problem8

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Problem8 struct{}

func (p *Problem8) Solve(data string, part int) int {
	registers := make(map[string]int)
	maxVal := math.MinInt
	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}
		processLine(&line, registers, &maxVal)
	}

	if part == 1 {
		return getLargestRegisterVal(registers)
	}
	return maxVal
}

func processLine(line *string, registers map[string]int, maxVal *int) {
	tokens := strings.Fields(strings.TrimSpace(*line))
	_, ok := registers[tokens[0]]
	if !ok {
		registers[tokens[0]] = 0
	}
	conditionValue, _ := strconv.Atoi(tokens[6])
	if isConditionTrue(registers[tokens[4]], tokens[5], conditionValue) {
		registerChangeVal, _ := strconv.Atoi(tokens[2])
		if tokens[1] == "inc" {
			registers[tokens[0]] += registerChangeVal
		} else {
			registers[tokens[0]] -= registerChangeVal
		}
		if registers[tokens[0]] > *maxVal {
			*maxVal = registers[tokens[0]]
		}
	}
}

func isConditionTrue(registerValue int, operator string, value int) bool {
	switch operator {
	case "<":
		return registerValue < value
	case ">":
		return registerValue > value
	case "<=":
		return registerValue <= value
	case ">=":
		return registerValue >= value
	case "==":
		return registerValue == value
	case "!=":
		return registerValue != value
	default:
		fmt.Println("Unknown operator", operator)
		return false
	}
}

func getLargestRegisterVal(registers map[string]int) int {
	largestVal := math.MinInt
	for _, val := range registers {
		if val > largestVal {
			largestVal = val
		}
	}
	return largestVal
}
