package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	var beginNode, endNode *ListNode = nil, head
	for node := head; node != nil; node = node.Next {
		k--
		if k == 0 {
			beginNode = node
		}
		if k < 0 {
			endNode = endNode.Next
		}
	}

	if beginNode != nil && endNode != nil {
		beginNode.Val, endNode.Val = endNode.Val, beginNode.Val
	}

	return head
}

func main() {
	head := &ListNode{4, &ListNode{1, &ListNode{3, &ListNode{2, &ListNode{7, nil}}}}}
	fmt.Println(swapNodes(head, 2))
}
