package utils

import (
	problems "AdventOfCode/Problems"
	"testing"
)

type TestData[T any] struct {
	Input    T
	Expected int
	Part     int
}

func CheckTestCases(t *testing.T, problem problems.Problem, testCases []TestData[string]) {
	for _, testCase := range testCases {
		actual := problem.Solve(testCase.Input, testCase.Part)
		if actual != testCase.Expected {
			t.Errorf("Wrong output for Input %q, Expected %d but got %d", testCase.Input, testCase.Expected, actual)
		}
	}
}
