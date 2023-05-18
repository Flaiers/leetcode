package main

import (
	"fmt"
	"sort"
)

func intersection(nums [][]int) []int {
	numsMap := make(map[int]int)
	for i := range nums {
		for j := range nums[i] {
			num := nums[i][j]
			numsMap[num]++
		}
	}

	var result []int
	for k, v := range numsMap {
		if v == len(nums) {
			result = append(result, k)
		}
	}

	sort.Ints(result)
	return result
}

func main() {
	fmt.Println(intersection([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}}))
}
