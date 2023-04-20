package main

import "fmt"

func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		ans[i] = ans[i>>1] + i%2
	}

	return ans
}

func main() {
	fmt.Println(countBits(7))
}
