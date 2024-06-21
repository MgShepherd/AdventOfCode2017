package main

import (
	problem1 "AdventOfCode/Problems/Problem1"
	utils "AdventOfCode/Utils"
	"fmt"
)

func main() {
	problem := problem1.Problem1{}
	fmt.Println(problem.Solve(utils.ReadProblemFile(1), 2))
}
