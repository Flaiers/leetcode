package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, next(head)

	for fast != nil {
		slow, fast = next(slow), next(next(fast))
	}

	return slow
}

func next(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}

	return nil
}

func main() {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(middleNode(l))
}
