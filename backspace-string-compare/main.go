package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	return removeBackspace(s) == removeBackspace(t)
}

func removeBackspace(s string) string {
	result := ""
	backspaceCount := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			backspaceCount++
			continue
		}

		if backspaceCount > 0 {
			backspaceCount--
			continue
		}

		result += string(s[i])
	}

	return result
}

func main() {
	fmt.Println(backspaceCompare("xywrrmp", "xywrrmu#p"))
}
