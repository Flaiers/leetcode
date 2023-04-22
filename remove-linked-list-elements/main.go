package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	root := &ListNode{Next: head}
	node := root
	for next(node) != nil {
		nextNode := next(node)
		if nextNode.Val == val {
			node.Next = next(nextNode)
		} else {
			node = nextNode
		}
	}

	return next(root)
}

func next(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}

	return nil
}

func main() {
	head := &ListNode{7, &ListNode{7, &ListNode{7, &ListNode{7, nil}}}}
	fmt.Println(removeElements(head, 7))
}
