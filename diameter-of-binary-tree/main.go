package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	answer := 0
	helper(root, &answer)
	return answer
}

func helper(node *TreeNode, answer *int) int {
	if node == nil {
		return 0
	}

	maxLeftDepth := helper(node.Left, answer)
	maxRightDepth := helper(node.Right, answer)

	*answer = max(*answer, maxLeftDepth+maxRightDepth)
	return max(maxLeftDepth, maxRightDepth) + 1
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
	fmt.Println(diameterOfBinaryTree(root))
}
