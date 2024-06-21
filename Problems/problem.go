package problems

type Problem interface {
	Solve(data string, part int) int
}
