package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := strconv.Itoa(root.Val)
	if root.Left != nil || root.Right != nil {
		result += "(" + tree2str(root.Left) + ")"
	}
	if root.Right != nil {
		result += "(" + tree2str(root.Right) + ")"
	}

	return result
}

func main() {
	fmt.Println(tree2str(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}))
}
