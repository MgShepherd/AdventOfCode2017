package problem5

import (
	utils "AdventOfCode/Utils"
	"testing"
)

func TestSolve(t *testing.T) {
	problem := &Problem5{}
	testCases := []utils.TestData{
		{Part: 1, Input: "0\n3\n0\n1\n-3", Expected: 5},
		{Part: 1, Input: "0\n3\n0\n1\n-3\n", Expected: 5},
		{Part: 1, Input: "0\n3 \n0\n1\n-3\n", Expected: 5},
		{Part: 2, Input: "0\n3 \n0\n1\n-3\n", Expected: 10},
	}

	utils.CheckTestCases(t, problem, testCases)
}
