package main

import "fmt"

func maxSubArray(nums []int) int {
	maxSum, curSum := nums[0], 0
	for _, num := range nums {
		curSum = max(num, curSum+num)
		maxSum = max(maxSum, curSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
