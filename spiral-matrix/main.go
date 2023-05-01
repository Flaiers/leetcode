package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var (
		n = len(matrix)
		m = len(matrix[0])

		up    = 0
		left  = 0
		down  = n - 1
		right = m - 1

		answer = make([]int, 0, n*m)
	)

	for len(answer) < cap(answer) {
		for i := left; i <= right && len(answer) < cap(answer); i++ {
			answer = append(answer, matrix[up][i])
		}
		up++

		for i := up; i <= down && len(answer) < cap(answer); i++ {
			answer = append(answer, matrix[i][right])
		}
		right--

		for i := right; i >= left && len(answer) < cap(answer); i-- {
			answer = append(answer, matrix[down][i])
		}
		down--

		for i := down; i >= up && len(answer) < cap(answer); i-- {
			answer = append(answer, matrix[i][left])
		}
		left++
	}

	return answer
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
