package main

import "fmt"

func reverseWords(s string) string {
	words := []byte(s)
	for i, j := 0, 0; j <= len(words); j++ {
		if j == len(words) || words[j] == ' ' {
			reverse(words[i:j])
			i = j + 1
		}
	}

	return string(words)
}

func reverse(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	fmt.Println(reverseWords("dnim eht pag"))
}
