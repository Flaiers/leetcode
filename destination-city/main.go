package main

import "fmt"

func destCity(paths [][]string) string {
	srcs := make(map[string]struct{}, len(paths))
	for _, path := range paths {
		srcs[path[0]] = struct{}{}
	}

	for _, path := range paths {
		if _, ok := srcs[path[1]]; !ok {
			return path[1]
		}
	}

	return ""
}

func main() {
	fmt.Println(destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}))
}
