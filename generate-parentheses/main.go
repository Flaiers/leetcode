package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string
	helper(n, 0, 0, "", &result)
	return result
}

func helper(n, opened, closed int, s string, result *[]string) {
	if len(s) == 2*n {
		*result = append(*result, s)
		return
	}

	if opened < n {
		helper(n, opened+1, closed, s+"(", result)
	}

	if closed < opened {
		helper(n, opened, closed+1, s+")", result)
	}
}

func main() {
	fmt.Println(generateParenthesis(5))
}
