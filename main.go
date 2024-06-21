package main

import (
	problem2 "AdventOfCode/Problems/Problem2"
	utils "AdventOfCode/Utils"
	"fmt"
)

func main() {
	problem := problem2.Problem2{}
	fmt.Println(problem.Solve(utils.ReadProblemFile(2), 2))
}
