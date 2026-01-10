package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	// 取到前后部分
	firstHalfEnd := endOfFirstHalf(head)
	// 反转后半部分
	secondHalfStart := reverseList(firstHalfEnd.Next)
	ptr1 := head
	ptr2 := secondHalfStart
	// 比对
	result := true
	for ptr2 != nil {
		if ptr1.Val != ptr2.Val {
			result = false
			break
		}
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	// 还原链表
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	result := isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}})
	fmt.Println(result)
}
