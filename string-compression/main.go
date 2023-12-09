package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	var prev byte
	var count int
	var answer int
	for _, c := range append(chars, ' ') {
		if c == prev {
			count++
			continue
		}

		if prev > 0 {
			chars[answer] = prev
			answer++
		}

		if count > 1 {
			for _, r := range strconv.Itoa(count) {
				chars[answer] = byte(r)
				answer++
			}
		}

		prev = c
		count = 1
	}

	return answer
}

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'a', 'a', 'b', 'b', 'b', 'b', 'c', 'c', 'c', 'c', 'd', 'd', 'd', 'd'}))
}
