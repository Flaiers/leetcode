package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	answer := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			answer = append(answer, "FizzBuzz")
		case i%3 == 0:
			answer = append(answer, "Fizz")
		case i%5 == 0:
			answer = append(answer, "Buzz")
		default:
			answer = append(answer, strconv.Itoa(i))
		}
	}

	return answer
}

func main() {
	fmt.Println(fizzBuzz(5))
}
