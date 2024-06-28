package main

import (
	problem8 "AdventOfCode/Problems/Problem8"
	utils "AdventOfCode/Utils"
	"fmt"
)

func main() {
	problem := problem8.Problem8{}
	fmt.Println(problem.Solve(utils.ReadProblemFile(8), 2))
}
