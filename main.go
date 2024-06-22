package main

import (
	problem4 "AdventOfCode/Problems/Problem4"
	utils "AdventOfCode/Utils"
	"fmt"
)

func main() {
	problem := problem4.Problem4{}
	fmt.Println(problem.Solve(utils.ReadProblemFile(4), 2))
}
