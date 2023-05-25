package main

func hIndex(citations []int) int {
	n := len(citations)
	count := make([]int, n+1)
	for _, citation := range citations {
		if citation > n {
			count[n]++
		} else {
			count[citation]++
		}
	}

	pos := 0
	for i := 0; i <= n; i++ {
		for j := 0; j < count[i]; j++ {
			citations[pos] = i
			pos++
		}
	}

	result := 1
	for i := n - 1; i >= 0; i-- {
		if citations[i] < result {
			return result - 1
		}
		result++
	}

	return n
}
