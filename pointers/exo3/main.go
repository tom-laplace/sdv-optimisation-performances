package main

import (
	"fmt"
	"sort"
)

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	res := mergeTwoLists(list1, list2)
	printList(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var temp []int
	var answer ListNode

	if list1 == nil && list2 == nil {
		return nil
	}

	for list1 != nil {
		temp = append(temp, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		temp = append(temp, list2.Val)
		list2 = list2.Next
	}

	sort.Ints(temp)

	var dummy *ListNode = &answer

	for i := range temp {
		dummy.Val = temp[i]

		if i < len(temp)-1 {
			dummy.Next = &ListNode{
				Val:  0,
				Next: nil,
			}

			dummy = dummy.Next
		}

	}

	return &answer
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}
