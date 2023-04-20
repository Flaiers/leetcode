package main

import "fmt"

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	return NumArray{prefixSum: prefixSum}
}

func (n *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return n.prefixSum[right]
	}

	return n.prefixSum[right] - n.prefixSum[left-1]
}

func main() {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(numArray.SumRange(0, 2))
	fmt.Println(numArray.SumRange(2, 5))
	fmt.Println(numArray.SumRange(0, 5))
}
