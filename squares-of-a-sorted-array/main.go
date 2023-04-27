package main

import "fmt"

func sortedSquares(nums []int) []int {
	lenNums := len(nums)
	answer := make([]int, lenNums)
	for i, j, k := 0, lenNums-1, lenNums-1; k >= 0; k-- {
		iSquare := nums[i] * nums[i]
		jSquare := nums[j] * nums[j]
		if iSquare > jSquare {
			answer[k] = iSquare
			i++
		} else {
			answer[k] = jSquare
			j--
		}
	}

	return answer
}

func main() {
	fmt.Println(sortedSquares([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}))
}
