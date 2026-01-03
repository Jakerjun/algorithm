package tool

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}
