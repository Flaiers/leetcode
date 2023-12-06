package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	store := make([][]int, n)
	for _, edge := range edges {
		store[edge[0]] = append(store[edge[0]], edge[1])
		store[edge[1]] = append(store[edge[1]], edge[0])
	}

	visit := make(map[int]bool, n)
	stack := make([]int, 0, n)

	stack = append(stack, source)
	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop == destination {
			return true
		}

		if !visit[pop] {
			visit[pop] = true
			stack = append(stack, store[pop]...)
		}
	}

	return false
}

func main() {}
