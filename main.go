package main

import (
	problem6 "AdventOfCode/Problems/Problem6"
	utils "AdventOfCode/Utils"
	"fmt"
)

func main() {
	problem := problem6.Problem6{}
	fmt.Println(problem.Solve(utils.ReadProblemFile(6), 2))
}
