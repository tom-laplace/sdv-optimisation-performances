package main

import (
	"testing"
)

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "Fusion de deux listes non vides",
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "Première liste vide",
			list1:    []int{},
			list2:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Deuxième liste vide",
			list1:    []int{1, 2, 3},
			list2:    []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Deux listes vides",
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := createList(tt.list1)
			l2 := createList(tt.list2)
			result := mergeTwoLists(l1, l2)
			got := listToSlice(result)

			if len(got) != len(tt.expected) {
				t.Errorf("longueur différente: got %v, expected %v", got, tt.expected)
				return
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("différence à l'index %d: got %v, expected %v", i, got, tt.expected)
					return
				}
			}
		})
	}
}
