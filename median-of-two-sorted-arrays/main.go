package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = merge(nums1, len(nums1), nums2, len(nums2))
	if len(nums1)%2 == 1 {
		return float64(nums1[len(nums1)/2])
	}

	return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	nums1 = append(nums1, make([]int, len(nums2))...)

	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}

	return nums1
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 8, 10}, []int{2, 4, 7, 9, 12}))
}
