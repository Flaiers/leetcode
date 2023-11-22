package main

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])

	dMatrix := make([][]int, m+n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dMatrix[i+j] = append(dMatrix[i+j], mat[i][j])
		}
	}

	nums := make([]int, 0, m*n)
	for i := 0; i < len(dMatrix); i++ {
		if i%2 == 0 {
			reverse(dMatrix[i])
		}
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

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
