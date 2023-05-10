package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x *= sign
	}

	result := 0
	for x > 0 {
		result = result*10 + x%10
		if result >= math.MaxInt32 {
			return 0
		}
		x /= 10
	}

	return sign * result
}

func main() {
	fmt.Println(reverse(123456789))
}
