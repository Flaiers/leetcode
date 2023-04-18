package main

import "fmt"

func dfs(grid [][]byte, i, j int) {
	grid[i][j] = '2'
	if i > 0 && grid[i-1][j] == '1' {
		dfs(grid, i-1, j)
	}
	if i < len(grid)-1 && grid[i+1][j] == '1' {
		dfs(grid, i+1, j)
	}
	if j > 0 && grid[i][j-1] == '1' {
		dfs(grid, i, j-1)
	}
	if j < len(grid[i])-1 && grid[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
}

func numIslands(grid [][]byte) int {
	var count int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
}
