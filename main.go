package main

import (
	problem9 "AdventOfCode/Problems/Problem9"
	utils "AdventOfCode/Utils"
	"fmt"
)

func main() {
	problem := problem9.Problem9{}
	fmt.Println(problem.Solve(utils.ReadProblemFile(9), 2))
}
