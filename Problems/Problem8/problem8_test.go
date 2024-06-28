package problem8

import (
	utils "AdventOfCode/Utils"
	"testing"
)

func TestSolve(t *testing.T) {
	problem := &Problem8{}
	testCases := []utils.TestData[int]{
		{Part: 1, Input: `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`, Expected: 1},
		{Part: 1, Input: `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1     
c inc -20 if c == 10`, Expected: 1},
		{Part: 1, Input: `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10
`, Expected: 1},
		{Part: 2, Input: `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`, Expected: 10},
	}

	utils.CheckTestCases(t, problem, testCases)
}
