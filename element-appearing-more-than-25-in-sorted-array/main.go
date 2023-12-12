package main

import "fmt"

func findSpecialInteger(arr []int) int {
	for i, j := 0, len(arr)/4; i < len(arr)-j; i++ {
		if arr[i] == arr[i+j] {
			return arr[i]
		}
	}

	return -1
}

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}
