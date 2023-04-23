package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head

	for node != nil && next(node) != nil {
		nextNode := next(node)
		if node.Val == nextNode.Val {
			node.Next = next(nextNode)
		} else {
			node = nextNode
		}
	}

	return head
}

func next(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}

	return nil
}

func main() {
	head := &ListNode{7, &ListNode{7, &ListNode{7, &ListNode{7, nil}}}}
	fmt.Println(deleteDuplicates(head))
}
