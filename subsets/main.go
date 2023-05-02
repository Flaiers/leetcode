package main

func subsets(nums []int) [][]int {
	var answer [][]int
	helper(&answer, nums, nil, 0)
	return answer
}

func helper(answer *[][]int, nums, subset []int, idx int) {
	if idx == len(nums) {
		subsetCopy := make([]int, len(subset))
		copy(subsetCopy, subset)
		*answer = append(*answer, subsetCopy)
		return
	}

	helper(answer, nums, subset, idx+1)
	subset = append(subset, nums[idx])
	helper(answer, nums, subset, idx+1)
	subset = subset[:len(subset)-1]
}
