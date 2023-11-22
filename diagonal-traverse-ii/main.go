package main

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	var (
		count int
		m     = len(mat)
		n     = len(mat[0])
	)
	for _, row := range mat {
		count += len(row)
		n = max(n, len(row))
	}

	dMatrix := make([][]int, m+n)
	for i := 0; i < m; i++ {
		for j := 0; j < len(mat[i]); j++ {
			dMatrix[i+j] = append(dMatrix[i+j], mat[i][j])
		}
	}

	nums := make([]int, 0, count)
	for i := 0; i < len(dMatrix); i++ {
		reverse(dMatrix[i])
		nums = append(nums, dMatrix[i]...)
	}

	return nums
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		j := len(nums) - i - 1
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}))
}
