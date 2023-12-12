package main

import "fmt"

func maxPathSum(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	minNum := nums[0]
	maxNum := nums[0]
	minNums := make([]float64, 4)
	maxNums := make([]float64, 4)

	for i := 1; i < len(nums); i++ {
		minNums[0] = minNum - nums[i]
		maxNums[0] = maxNum - nums[i]

		minNums[1] = minNum + nums[i]
		maxNums[1] = maxNum + nums[i]

		minNums[2] = minNum * nums[i]
		maxNums[2] = maxNum * nums[i]

		if nums[i] > 0 {
			minNums[3] = minNum / nums[i]
			maxNums[3] = maxNum / nums[i]
		}

		minNum = min(min(minNums...), min(maxNums...))
		maxNum = max(max(minNums...), max(maxNums...))
	}

	return maxNum
}

func min(nums ...float64) float64 {
	minimum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > minimum {
			minimum = nums[i]
		}
	}

	return minimum
}

func max(nums ...float64) float64 {
	maximum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maximum {
			maximum = nums[i]
		}
	}

	return maximum
}

func main() {
	fmt.Println(maxPathSum([]float64{-1.2, 0.12, -7.1, 0, 12.4, 0.34, 2.0}))
}
