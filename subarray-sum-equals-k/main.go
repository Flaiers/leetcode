package main

import "fmt"

func subarraySum(nums []int, k int) int {
	sum := 0
	result := 0
	sumCount := map[int]int{0: 1}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		result += sumCount[sum-k]
		sumCount[sum]++
	}

	return result
}

func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
