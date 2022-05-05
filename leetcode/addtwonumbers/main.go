package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{
		Val: 0,
	}
	curr := head

	for l1 != nil || l2 != nil {
		val1 := func(node *ListNode) int {
			if node == nil {
				return 0
			}
			return node.Val
		}(l1)
		val2 := func(node *ListNode) int {
			if node == nil {
				return 0
			}
			return node.Val
		}(l2)
		sum := val1 + val2 + carry
		carry = sum / 10
		curr.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}

	return head.Next
}
