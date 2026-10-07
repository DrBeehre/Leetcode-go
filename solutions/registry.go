package solutions

import "sort"

type Solution struct {
	ID         int
	Title      string
	Difficulty string
	Run        func()
}

var registry = map[int]Solution{}

func RegisterSolution(solution Solution) {
	registry[solution.ID] = solution
}

func GetSolution(id int) (Solution, bool) {
	solution, ok := registry[id]
	return solution, ok
}

func GetAllSolutions() []Solution {
	solutions := make([]Solution, 0, len(registry))
	for _, solution := range registry {
		solutions = append(solutions, solution)
	}
	sort.Slice(solutions, func(i, j int) bool {
		return solutions[i].ID < solutions[j].ID
	})
	return solutions
}
