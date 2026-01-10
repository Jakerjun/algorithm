package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	cur := head
	next := &ListNode{}
	for cur != nil && cur.Next != nil {
		// 取next
		tempNext := cur.Next
		// 改next
		cur.Next = next
		// 存 next
		next = tempNext
		cur = tempNext

	}

	if next != nil {
		cur.Next = next
	}

	return cur
}

func main() {
	result := reverseList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}})
	printListNode(result)
}
