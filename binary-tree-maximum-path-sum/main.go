package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	answer := -1 << 63
	helper(root, &answer)
	return answer
}

func helper(node *TreeNode, answer *int) int {
	if node == nil {
		return 0
	}

	maxLeftPath := helper(node.Left, answer)
	maxRightPath := helper(node.Right, answer)

	*answer = max(*answer, maxLeftPath+maxRightPath+node.Val)
	return max(max(maxLeftPath+node.Val, maxRightPath+node.Val), 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
	fmt.Println(maxPathSum(root))
}
