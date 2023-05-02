package main

func longestConsecutive(nums []int) int {
	uniqNums := make(map[int]struct{})
	for _, num := range nums {
		uniqNums[num] = struct{}{}
	}

	maxLength := 0
	for num := range uniqNums {
		if has(uniqNums, num-1) {
			continue
		}

		currentLength := 1
		for has(uniqNums, num+currentLength) {
			currentLength++
		}

		maxLength = max(maxLength, currentLength)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func has(m map[int]struct{}, key int) bool {
	_, ok := m[key]
	return ok
}
