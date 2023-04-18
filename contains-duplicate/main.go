package main

import "fmt"

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			return true
		}
		numsMap[num] = struct{}{}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 2, 3, 4, 5}))
}
