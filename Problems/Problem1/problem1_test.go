package problem1

import "testing"

type testData struct {
	input    string
	expected int
	part     int
}

func TestSolve(t *testing.T) {
	problem := Problem1{}
	testCases := []testData{
		{part: 1, input: "1122", expected: 3},
		{part: 1, input: "1111", expected: 4},
		{part: 1, input: "1234", expected: 0},
		{part: 1, input: "91212129", expected: 9},
		{part: 1, input: "91212129 ", expected: 9},
		{part: 1, input: "", expected: 0},
		{part: 2, input: "1212", expected: 6},
		{part: 2, input: "1221", expected: 0},
		{part: 2, input: "123425", expected: 4},
		{part: 2, input: "123123", expected: 12},
		{part: 2, input: "123123 ", expected: 12},
		{part: 2, input: "12131415", expected: 4},
		{part: 2, input: "", expected: 0},
	}

	for _, testCase := range testCases {
		actual := problem.Solve(testCase.input, testCase.part)
		if actual != testCase.expected {
			t.Errorf("Wrong output for input %q, expected %d but got %d", testCase.input, testCase.expected, actual)
		}
	}
}
