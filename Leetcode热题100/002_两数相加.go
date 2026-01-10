package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 有为0的
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}
	if l2.Val == 0 && l2.Next == nil {
		return l1
	}
	dynamic := new(ListNode)
	ptr := dynamic
	sum := 0
	// 同位相加
	for l1 != nil && l2 != nil {
		sum += l1.Val + l2.Val
		ptr.Next = &ListNode{Val: sum % 10}
		ptr = ptr.Next
		l1 = l1.Next
		l2 = l2.Next
		if sum >= 10 {
			sum = 1
		} else {
			sum = 0
		}
	}

	// 此刻 肯定有遍历完了的 那么分类讨论 如果是两个都遍历完了 那么就剩下最后一个进位信号了
	for l1 != nil { // 有一个还没有遍历 那么只需要加上进位信号
		sum += l1.Val
		ptr.Next = &ListNode{Val: sum % 10}
		ptr = ptr.Next
		l1 = l1.Next
		if sum >= 10 {
			sum = 1
		} else {
			sum = 0
		}
	}
	for l2 != nil {
		sum += l2.Val
		ptr.Next = &ListNode{Val: sum % 10}
		ptr = ptr.Next
		l2 = l2.Next
		if sum >= 10 {
			sum = 1
		} else {
			sum = 0
		}
	}

	if sum == 1 {
		ptr.Next = &ListNode{Val: 1}
	}

	return dynamic.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	result := addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}})
	// result是一个链表值
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}
