package main

import "fmt"

func reverseWords(words []byte) []byte {
	reverse(words)
	for i, j := 0, 0; j <= len(words); j++ {
		if j == len(words) || words[j] == ' ' {
			reverse(words[i:j])
			i = j + 1
		}
	}

	return words
}

func reverse(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	fmt.Printf("%s\n", reverseWords([]byte{'g', 'a', 'p', ' ', 't', 'h', 'e', ' ', 'm', 'i', 'n', 'd'}))
}
