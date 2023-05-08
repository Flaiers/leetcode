package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])

	result := make([][]int, n)
	for r := range result {
		result[r] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
