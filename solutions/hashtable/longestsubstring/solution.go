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
	charMap := make(map[rune]int) // Seen: pos

	substring := ""
	longestSubstringLength := 0

	for pos, char := range []rune(s) {
		fmt.Printf("pos: %v, char: %c\n", pos, char)

		// I KNOW! I need to create a new substring starting from the position of the last seen position!

		seenPos, found := charMap[char]
		if found {
			fmt.Printf("seen char %c before at position %v\n", char, seenPos)
			substring = string(char)
			// NEED TO USE THE SEEN POSITION SOMEWHERE

		} else {
			substring += string(char)
		}

		if len(substring) > longestSubstringLength {
			longestSubstringLength = len(substring)
		}

		charMap[char] = pos
	}

	return longestSubstringLength
}
