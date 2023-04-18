package main

import "fmt"

func permuteRecursive(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		*result = append(*result, append([]int{}, nums...))
		return
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		permuteRecursive(nums, start+1, result)
		nums[i], nums[start] = nums[start], nums[i]
	}
}

func permute(nums []int) [][]int {
	var result [][]int
	permuteRecursive(nums, 0, &result)
	return result
}

func main() {
	fmt.Println(permute([]int{0, 1}))
}
