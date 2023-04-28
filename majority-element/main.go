package main

import "fmt"

func majorityElement(nums []int) int {
	count := 1
	answer := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == answer {
			count++
		} else {
			count--
		}

		if count == 0 {
			count = 1
			answer = nums[i]
		}
	}

	return answer
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
