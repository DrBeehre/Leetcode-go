package longestsubstring

import (
	"fmt"

	"github.com/DrBeehre/Leetcode-go/solutions"
)

func init() {
	solutions.RegisterSolution(solutions.Solution{
		ID:         3,
		Title:      "Longest Substring Without Repeating Characters",
		Difficulty: "Medium",
		Run: func() {
			fmt.Println("Longest Substring is a work in progress...")
			fmt.Printf("%d", lengthOfLongestSubstring("abcabcbb"))
		},
	})
}

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	startPos := 0
	longestSubStringLength := 0

	for pos, char := range []rune(s) {

		seenPos, found := charMap[char]
		if found && seenPos >= startPos {
			startPos = seenPos + 1
		}

		length := pos - startPos + 1
		if length > longestSubStringLength {
			longestSubStringLength = length
		}

		charMap[char] = pos
	}

	return longestSubStringLength
}
