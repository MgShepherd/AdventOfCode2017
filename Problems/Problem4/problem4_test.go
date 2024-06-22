package problem4

import (
	utils "AdventOfCode/Utils"
	"testing"
)

func TestSolve(t *testing.T) {
	problem := &Problem4{}
	testCases := []utils.TestData{
		{Part: 1, Input: "aa bb cc dd ee", Expected: 1},
		{Part: 1, Input: "aa bb cc dd aa", Expected: 0},
		{Part: 1, Input: "aa bb cc dd aaa", Expected: 1},
		{Part: 1, Input: "aa bb cc dd aaa\naa bb cc dd ee", Expected: 2},
		{Part: 1, Input: "aa bb cc dd aaa ", Expected: 1},
		{Part: 1, Input: "aa bb cc dd aaa \n\n", Expected: 1},
		{Part: 2, Input: "abcde fghij ", Expected: 1},
		{Part: 2, Input: "abcde xyz ecdab ", Expected: 0},
		{Part: 2, Input: "a ab abc abd abf abj", Expected: 1},
		{Part: 2, Input: "iiii oiii ooii oooi oooo\noiii ioii iioi iiio", Expected: 1},
	}

	utils.CheckTestCases(t, problem, testCases)
}
