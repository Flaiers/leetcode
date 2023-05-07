package main

import "fmt"

func sortColors(nums []int) {
	red, white, blue := 0, 0, 0
	for _, num := range nums {
		switch num {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		}
	}

	i := 0
	for red != 0 {
		nums[i] = 0
		red--
		i++
	}
	for white != 0 {
		nums[i] = 1
		white--
		i++
	}
	for blue != 0 {
		nums[i] = 2
		blue--
		i++
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
