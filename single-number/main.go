package main

import "fmt"

func singleNumber(nums []int) int {
	var mask int
	for _, num := range nums {
		mask ^= num
	}

	return mask
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
