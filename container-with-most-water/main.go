package main

import "fmt"

func maxArea(height []int) int {
	i, j := 0, len(height)-1

	var maximum int
	for i < j {
		current := min(height[i], height[j]) * (j - i)
		maximum = max(maximum, current)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
