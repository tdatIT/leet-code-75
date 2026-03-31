package main

import "fmt"

func main() {
	linkedlist := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	fmt.Println(reverseList(linkedlist))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	for head != nil {
		next := head.Next
		current := head
		current.Next = pre
		pre = current

		head = next
	}

	return pre
}
