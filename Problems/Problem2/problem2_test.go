package problem2

import (
	utils "AdventOfCode/Utils"
	"testing"
)

func TestSolve(t *testing.T) {
	problem := &Problem2{}
	testCases := []utils.TestData[string]{
		{Part: 1, Input: "5 1 9 5", Expected: 8},
		{Part: 1, Input: "7 5 3", Expected: 4},
		{Part: 1, Input: "7 5 3  ", Expected: 4},
		{Part: 1, Input: "2\t4 6 8", Expected: 6},
		{Part: 1, Input: "5 1 9 5\n2 4 6 8", Expected: 14},
		{Part: 2, Input: "5 9 2\t8", Expected: 4},
		{Part: 2, Input: "9 4 7 3", Expected: 3},
		{Part: 2, Input: "5 9 2\t8\n3 8 6 5", Expected: 6},
	}

	utils.CheckTestCases(t, problem, testCases)
}
