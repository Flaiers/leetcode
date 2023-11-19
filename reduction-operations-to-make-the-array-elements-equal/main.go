package main

import "fmt"

func reductionOperations(nums []int) int {
	var largest int
	for _, num := range nums {
		if num > largest {
			largest = num
		}
	}

	countNums := make([]int, largest+1)
	for _, num := range nums {
		countNums[num]++
	}

	var count int
	var operations int
	for i := largest; i > 0; i-- {
		if countNums[i] > 0 {
			count += countNums[i]
			operations += count - countNums[i]
		}
	}

	return operations
}

func main() {
	fmt.Println(reductionOperations([]int{1, 1, 2, 2, 3}))
}
