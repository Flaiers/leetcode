package main

import "fmt"

func isMonotonic(input []int) bool {
	increasing, decreasing := true, true
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			decreasing = false
		}
		if input[i] < input[i-1] {
			increasing = false
		}
	}
	return increasing || decreasing
}

func main() {
	input := []int{}
	fmt.Println(isMonotonic(input))
}
