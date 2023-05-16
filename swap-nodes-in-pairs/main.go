package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		head, head.Next, head.Next.Next = head.Next, swapPairs(head.Next.Next), head
	}

	return head
}

func main() {
	head := &ListNode{4, &ListNode{2, &ListNode{3, &ListNode{1, nil}}}}
	fmt.Println(swapPairs(head))
}
