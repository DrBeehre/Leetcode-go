package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DrBeehre/Leetcode-go/solutions"
)

func Execute() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		if len(solutions.GetAllSolutions()) == 0 {
			fmt.Println("No solutions registered.")
		}
		for _, solution := range solutions.GetAllSolutions() {
			fmt.Printf("ID: %d, Title: %s, Difficulty: %s\n", solution.ID, solution.Title, solution.Difficulty)
		}
	case "run":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Error: 'run' command requires a solution ID\n")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: invalid solution ID: %s\n. Id must be a number", os.Args[2])
			os.Exit(1)
		}
		solution, ok := solutions.GetSolution(id)
		if !ok {
			fmt.Fprintf(os.Stderr, "Error: no solution found with ID %d\n", id)
			os.Exit(1)
		}
		solution.Run()
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command '%s'\n", os.Args[1])
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("usage: leetcode <list|run> [soltion_id]")
}
