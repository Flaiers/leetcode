package main

import "fmt"

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	multiplier := 1
	for i := 0; i < len(nums); i++ {
		answer[i] = multiplier
		multiplier *= nums[i]
	}

	multiplier = 1
	for i := len(answer) - 1; i >= 0; i-- {
		answer[i] *= multiplier
		multiplier *= nums[i]
	}

	return answer
}

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
