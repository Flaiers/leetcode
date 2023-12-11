package main

import "fmt"

func reverseWords(s string) string {
	words := reverse([]byte(s))
	var result string
	for i, j := 0, 0; j <= len(words); j++ {
		var word string
		if j == len(words) || words[j] == ' ' {
			word = string(reverse(words[i:j]))
			i = j + 1
		}

		if result != "" && word != "" {
			result += " "
		}
		result += word
	}

	return result
}

func reverse(s []byte) []byte {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func main() {
	fmt.Println(reverseWords("  gap  the  mind  "))
}
