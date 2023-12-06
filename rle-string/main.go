package main

import (
	"fmt"
	"strconv"
)

func encode(s string) string {
	var prev rune
	var count int
	var answer string
	for _, r := range s + " " {
		if r == prev {
			count++
			continue
		}

		answer += string(prev)
		if count > 1 {
			answer += strconv.Itoa(count)
		}

		prev = r
		count = 1
	}

	return answer
}

func main() {
	fmt.Println(encode("aaaabbbbccccdddd"))
}
