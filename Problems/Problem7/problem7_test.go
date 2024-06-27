package problem7

import (
	utils "AdventOfCode/Utils"
	"testing"
)

func TestSolve(t *testing.T) {
	problem := &Problem7{}
	testCases := []utils.TestData[string]{
		{Part: 1, Input: `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`, Expected: "tknk"},
		{Part: 1, Input: `pbga (66)
xhth (57)
    ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)\n`, Expected: "tknk"},
		{Part: 2, Input: `pbga (66)
xhth (57)
    ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)\n`, Expected: "60"},
	}

	utils.CheckTestCases(t, problem, testCases)
}
