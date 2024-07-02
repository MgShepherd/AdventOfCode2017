package problem9

import (
	utils "AdventOfCode/Utils"
	"testing"
)

func TestSolve(t *testing.T) {
	problem := &Problem9{}
	testCases := []utils.TestData[int]{
		{Part: 1, Input: "{}", Expected: 1},
		{Part: 1, Input: "{{{}}}", Expected: 6},
		{Part: 1, Input: "{{},{}}\n{{{},{},{{}}}}", Expected: 21},
		{Part: 1, Input: "{<a>,<a>,<a>,<a>}", Expected: 1},
		{Part: 1, Input: "{{<ab>},{<ab>},{<ab>},{<ab>}}", Expected: 9},
		{Part: 1, Input: "{{<!!>},{<!!>},{<!!>},{<!!>}}", Expected: 9},
		{Part: 1, Input: "{{<a!>},{<a!>},{<a!>},{<ab>}}", Expected: 3},
		{Part: 2, Input: "{<>}", Expected: 0},
		{Part: 2, Input: "{<random characters>}", Expected: 17},
		{Part: 2, Input: "{<<<<>}", Expected: 3},
		{Part: 2, Input: "{<{!>}>}", Expected: 2},
		{Part: 2, Input: "{<!!>}", Expected: 0},
		{Part: 2, Input: "{<{osi!a,<{i<a>}", Expected: 10},
	}
	utils.CheckTestCases(t, problem, testCases)
}
