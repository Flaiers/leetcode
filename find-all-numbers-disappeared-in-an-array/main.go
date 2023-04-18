package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		pos := nums[i] - 1
		if nums[i] != nums[pos] {
			nums[i], nums[pos] = nums[pos], nums[i]
		} else {
			i++
		}
	}

	miss := make([]int, 0)
	for idx, num := range nums {
		if num != idx+1 {
			miss = append(miss, idx+1)
		}
	}

	return miss
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
