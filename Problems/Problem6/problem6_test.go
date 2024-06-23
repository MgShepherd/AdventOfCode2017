package problem6

import (
	utils "AdventOfCode/Utils"
	"testing"
)

func TestSolve(t *testing.T) {
	problem := &Problem6{}
	testCases := []utils.TestData{
		{Part: 1, Input: "0 2 7 0", Expected: 5},
		{Part: 1, Input: "0\t 2\t7 0", Expected: 5},
		{Part: 2, Input: "0\t 2\t7 0", Expected: 4},
	}

	utils.CheckTestCases(t, problem, testCases)
}
