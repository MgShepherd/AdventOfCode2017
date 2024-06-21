package problem1

import (
	utils "AdventOfCode/Utils"
	"testing"
)

func TestSolve(t *testing.T) {
	problem := &Problem1{}
	testCases := []utils.TestData{
		{Part: 1, Input: "1122", Expected: 3},
		{Part: 1, Input: "1111", Expected: 4},
		{Part: 1, Input: "1234", Expected: 0},
		{Part: 1, Input: "91212129", Expected: 9},
		{Part: 1, Input: "91212129 ", Expected: 9},
		{Part: 1, Input: "", Expected: 0},
		{Part: 2, Input: "1212", Expected: 6},
		{Part: 2, Input: "1221", Expected: 0},
		{Part: 2, Input: "123425", Expected: 4},
		{Part: 2, Input: "123123", Expected: 12},
		{Part: 2, Input: "123123 ", Expected: 12},
		{Part: 2, Input: "12131415", Expected: 4},
		{Part: 2, Input: "", Expected: 0},
	}

	utils.CheckTestCases(t, problem, testCases)
}
