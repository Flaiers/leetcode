package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	down, left := 0, n-1

	for down < m && left >= 0 {
		if matrix[down][left] == target {
			return true
		} else if matrix[down][left] > target {
			left--
		} else if matrix[down][left] < target {
			down++
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 5))
}
