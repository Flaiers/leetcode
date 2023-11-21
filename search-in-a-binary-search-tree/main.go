package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > val:
		return searchBST(root.Left, val)
	case root.Val < val:
		return searchBST(root.Right, val)
	default:
		return root
	}
}

func main() {
	fmt.Println(searchBST(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}, 5))
}
