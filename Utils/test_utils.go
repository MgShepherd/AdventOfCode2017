package utils

import (
	problems "AdventOfCode/Problems"
	"testing"
)

type TestData[T problems.ProblemOutput] struct {
	Input    string
	Expected T
	Part     int
}

func CheckTestCases[T problems.ProblemOutput](t *testing.T, problem problems.Problem[T], testCases []TestData[T]) {
	for _, testCase := range testCases {
		actual := problem.Solve(testCase.Input, testCase.Part)
		if actual != testCase.Expected {
			t.Errorf("Wrong output for Input %v, Expected %v but got %v", testCase.Input, testCase.Expected, actual)
		}
	}
}
