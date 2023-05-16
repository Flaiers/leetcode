package main

import "fmt"

func generate(numRows int) [][]int {
	answer := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		answer[i] = make([]int, i+1)
		answer[i][0], answer[i][i] = 1, 1
		for j := 1; j < i; j++ {
			answer[i][j] = answer[i-1][j-1] + answer[i-1][j]
		}
	}

	return answer
}

func main() {
	fmt.Println(generate(5))
}
