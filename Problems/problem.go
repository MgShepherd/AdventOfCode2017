package problems

type ProblemOutput interface {
	int | string
}

type Problem[T ProblemOutput] interface {
	Solve(data string, part int) T
}
