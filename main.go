package main

import (
	problem5 "AdventOfCode/Problems/Problem5"
	utils "AdventOfCode/Utils"
	"fmt"
)

func main() {
	problem := problem5.Problem5{}
	fmt.Println(problem.Solve(utils.ReadProblemFile(5), 2))
}
