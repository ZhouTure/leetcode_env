package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(nums []int) *ListNode {
	dummy := &ListNode{} // 哑节点
	tail := dummy
	for _, v := range nums {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

func printList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val)
		if p.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			tail.Next = p1
			p1 = p1.Next
		} else {
			tail.Next = p2
			p2 = p2.Next
		}
		tail = tail.Next
	}
	if p1 != nil {
		tail.Next = p1
	} else {
		tail.Next = p2
	}

	return dummy.Next
}

func main1() {
	nums1 := []int{1, 2, 4}
	nums2 := []int{1, 3, 4}
	// nums := []int{1, 2, 2, 1}
	head1 := buildList(nums1)
	head2 := buildList(nums2)
	res := mergeTwoLists(head1, head2)
	fmt.Printf("res: ")
	printList(res)
}
