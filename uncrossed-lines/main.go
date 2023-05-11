package main

import "fmt"

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	prev := make([]int, l2+1)

	for i := 1; i <= l1; i++ {
		cur := make([]int, l2+1)
		for j := 1; j <= l2; j++ {
			if nums1[i-1] == nums2[j-1] {
				cur[j] = prev[j-1] + 1
			} else {
				cur[j] = max(cur[j-1], prev[j])
			}
		}
		prev = cur
	}

	return prev[l2]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}))
}
