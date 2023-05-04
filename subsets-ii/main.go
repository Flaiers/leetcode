package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var answer [][]int
	sort.Ints(nums)
	helper(&answer, nums, nil, 0)
	return answer
}

func helper(answer *[][]int, nums, subset []int, idx int) {
	*answer = append(*answer, append([]int{}, subset...))
	for i := idx; i < len(nums); i++ {
		if i == idx || nums[i] != nums[i-1] {
			helper(answer, nums, append(subset, nums[i]), i+1)
		}
	}
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 3, 2}))
}
