package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	numCount := make(map[int]int, len(nums))
	for _, num := range nums {
		numCount[num]++
	}

	countNums := make([][]int, len(nums)+1)
	for num, count := range numCount {
		countNums[count] = append(countNums[count], num)
	}

	var answer []int
	for i := len(countNums) - 1; i > 0; i-- {
		if len(countNums[i]) > 0 {
			answer = append(answer, countNums[i]...)
		}
	}

	return answer[:k]
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
