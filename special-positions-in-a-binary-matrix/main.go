package main

import "fmt"

func numSpecial(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	rowsSum := make([]int, n)
	columnsSum := make([]int, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rowsSum[j] += mat[i][j]
			columnsSum[i] += mat[i][j]
		}
	}

	var count int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && rowsSum[j] == 1 && columnsSum[i] == 1 {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(numSpecial([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
}
