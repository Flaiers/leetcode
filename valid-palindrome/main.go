package main

import "fmt"

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func toLower(r byte) byte {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlphanumeric(rune(s[i])) {
			i++
		}
		for i < j && !isAlphanumeric(rune(s[j])) {
			j--
		}
		if i < j && toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
