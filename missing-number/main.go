package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	return n*(n+1)/2 - sum(nums)
}

func sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
