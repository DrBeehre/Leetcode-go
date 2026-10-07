package twosum

import (
	"fmt"

	"github.com/DrBeehre/Leetcode-go/solutions"
)

func init() {
	solutions.RegisterSolution(solutions.Solution{
		ID:         1,
		Title:      "Two Sum",
		Difficulty: "Easy",
		Run: func() {
			fmt.Println("Running Two Sum solution...") // here I can call a new function that implements the solution for Two Sum
		},
	})
}
