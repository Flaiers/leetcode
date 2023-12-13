package main

import "fmt"

func maxProduct(nums []int) int {
	maxNum1, maxNum2 := nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		if nums[i] > maxNum1 {
			maxNum2 = max(maxNum1, maxNum2)
			maxNum1 = nums[i]
		} else if nums[i] > maxNum2 {
			maxNum1 = max(maxNum1, maxNum2)
			maxNum2 = nums[i]
		}
	}

	return (maxNum1 - 1) * (maxNum2 - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
}
