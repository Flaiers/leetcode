package main

import "fmt"

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for r := range result {
		result[r] = make([]int, n)
	}

	var (
		up    = 0
		left  = 0
		down  = n - 1
		right = n - 1

		number = 1
	)

	for number <= n*n {
		for i := left; i <= right; i++ {
			result[up][i] = number
			number++
		}
		up++

		for i := up; i <= down; i++ {
			result[i][right] = number
			number++
		}
		right--

		for i := right; i >= left; i-- {
			result[down][i] = number
			number++
		}
		down--

		for i := down; i >= up; i-- {
			result[i][left] = number
			number++
		}
		left++
	}

	return result
}

func main() {
	fmt.Println(generateMatrix(3))
}
