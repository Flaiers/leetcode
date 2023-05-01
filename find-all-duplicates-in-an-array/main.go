package main

import "fmt"

func findDuplicates(nums []int) []int {
	answer := make([]int, 0, len(nums))
	for _, num := range nums {
		absNum := abs(num)
		if nums[absNum-1] < 0 {
			answer = append(answer, absNum)
		}
		nums[absNum-1] *= -1
	}

	return answer
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
