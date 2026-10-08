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
			// fmt.Println("Running Two Sum solution...")
			fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
		},
	})
}

func twoSum(nums []int, target int) []int {
	var targetPair []int
	numMap := make(map[int]int)

	for pos, num := range nums {
		requiredTarget := target - num
		matchingPairPos, found := numMap[requiredTarget]
		if found {
			targetPair = []int{matchingPairPos, pos}
		}
		numMap[num] = pos
	}

	if len(targetPair) == 0 {
		return nil
	}

	return targetPair
}
