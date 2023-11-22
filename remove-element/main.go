package main

import "fmt"

func removeElement(nums []int, val int) int {
	var length int
	for _, num := range nums {
		if num != val {
			nums[length] = num
			length++
		}
	}

	return length
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
