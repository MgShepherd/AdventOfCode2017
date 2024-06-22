package problem3

import (
	utils "AdventOfCode/Utils"
	"testing"
)

func TestSolve(t *testing.T) {
	problem := &Problem3{}
	testCases := []utils.TestData{
		{Part: 1, Input: "1", Expected: 0},
		{Part: 1, Input: "12", Expected: 3},
		{Part: 1, Input: "23", Expected: 2},
		{Part: 1, Input: "1024", Expected: 31},
		{Part: 2, Input: "1", Expected: 2},
		{Part: 2, Input: "100", Expected: 122},
		{Part: 2, Input: "30", Expected: 54},
		{Part: 2, Input: "800", Expected: 806},
	}

	utils.CheckTestCases(t, problem, testCases)
}
