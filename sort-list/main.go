package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	node := head
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	head = node
	sort.Ints(values)
	for _, value := range values {
		head.Val = value
		head = head.Next
	}

	return node
}

func main() {
	head := &ListNode{4, &ListNode{1, &ListNode{3, &ListNode{2, nil}}}}
	fmt.Println(sortList(head))
}
