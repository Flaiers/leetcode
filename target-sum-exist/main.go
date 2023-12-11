package main

import "fmt"

func targetSumExist(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	return helper(nums, target, nums[0], 1)
}

func helper(nums []int, target, result, idx int) bool {
	if idx == len(nums) {
		return target == result
	}

	return helper(nums, target, result+nums[idx], idx+1) ||
		helper(nums, target, result-nums[idx], idx+1)
}

func main() {
	fmt.Println(targetSumExist([]int{9, 3, 7}, 13))
}
