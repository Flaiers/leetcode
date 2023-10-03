package main

import "fmt"

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return compareMap(stringToMap(s), stringToMap(t))
}

func stringToMap(s string) map[rune]int {
	m := make(map[rune]int, len(s))
	for _, r := range s {
		m[r]++
	}
	return m
}

func compareMap(s, t map[rune]int) bool {
	for sk, sv := range s {
		if tv, ok := t[sk]; !ok || sv != tv {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
