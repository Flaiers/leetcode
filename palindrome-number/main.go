package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == reverse(x)
}

func reverse(x int) int {
	result := 0
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}

	return result
}

func main() {
	fmt.Println(isPalindrome(1234321))
}
