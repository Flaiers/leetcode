package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	diff := 1<<63 - 1
	dfs(root, &diff)
	return diff
}

func dfs(root *TreeNode, diff *int) (int, int) {
	minValue, maxValue := root.Val, root.Val
	if root.Left != nil {
		minVal, maxVal := dfs(root.Left, diff)
		*diff = min(*diff, root.Val-maxVal)
		minValue = minVal
	}
	if root.Right != nil {
		minVal, maxVal := dfs(root.Right, diff)
		*diff = min(*diff, minVal-root.Val)
		maxValue = maxVal
	}

	return minValue, maxValue
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(getMinimumDifference(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
			Left: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}))
}
