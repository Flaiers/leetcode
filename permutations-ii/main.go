package main

import "fmt"

func permuteUnique(nums []int) [][]int {
	var result [][]int
	permuteRecursive(nums, 0, &result)
	return result
}

func permuteRecursive(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		*result = append(*result, append([]int{}, nums...))
		return
	}
	used := make(map[int]struct{})
	for i := start; i < len(nums); i++ {
		if _, ok := used[nums[i]]; ok {
			continue
		}
		nums[i], nums[start] = nums[start], nums[i]
		permuteRecursive(nums, start+1, result)
		nums[i], nums[start] = nums[start], nums[i]
		used[nums[i]] = struct{}{}
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}
