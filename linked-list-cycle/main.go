package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, next(head)

	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow, fast = next(slow), next(next(fast))
	}

	return false
}

func next(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}

	return nil
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = head
	fmt.Println(hasCycle(head))
}
