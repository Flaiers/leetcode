package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	node := reverseList(middleNode(head))

	for head != nil && node != nil {
		if head.Val != node.Val {
			return false
		}

		head = next(head)
		node = next(node)
	}

	return true
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, next(head)

	for fast != nil {
		slow, fast = next(slow), next(next(fast))
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func next(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}

	return nil
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	fmt.Println(isPalindrome(head))
}
