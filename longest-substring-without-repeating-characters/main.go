package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	start := 0
	result := 0
	passed := make(map[rune]int)

	for i, c := range s {
		if v, ok := passed[c]; ok && v >= start {
			start = v + 1
		}

		result = max(result, i-start+1)
		passed[c] = i
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
