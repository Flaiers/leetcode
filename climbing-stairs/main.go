package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	n1, n2 := 1, 2
	for i := 3; i < n+1; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}

func main() {
	fmt.Println(climbStairs(7))
}
