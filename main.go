package main

import (
	problem6 "AdventOfCode/Problems/Problem7"
	utils "AdventOfCode/Utils"
	"fmt"
)

func main() {
	problem := problem6.Problem7{}
	fmt.Println(problem.Solve(utils.ReadProblemFile(7), 2))
}
