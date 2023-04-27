package main

import "fmt"

func twoSum(nums []int, target int) []int {
	kv := make(map[int]int)
	for index, num := range nums {
		if other, ok := kv[num]; !ok {
			kv[target-num] = index
		} else {
			return []int{other, index}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 10, 1}, 3))
}
