package main

import "fmt"

func largestGoodInteger(num string) string {
	j := -1
	for i := 1; i < len(num)-1; i++ {
		if num[i-1] == num[i] && num[i] == num[i+1] {
			if j == -1 || gt(num[i], num[j]) {
				j = i
			}
		}
	}

	if j == -1 {
		return ""
	}

	return num[j-1 : j+2]
}

func gt(a, b byte) bool {
	if int(a)-'0' > int(b)-'0' {
		return true
	}
	return false
}

func main() {
	fmt.Println(largestGoodInteger("6777133339"))
}
