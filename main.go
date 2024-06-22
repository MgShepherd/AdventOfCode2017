package main

import (
	problem3 "AdventOfCode/Problems/Problem3"
	utils "AdventOfCode/Utils"
	"fmt"
)

func main() {
	problem := problem3.Problem3{}
	fmt.Println(problem.Solve(utils.ReadProblemFile(3), 2))
}
