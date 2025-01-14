package main

import "fmt"

type ListNode struct {
	Val  int32
	Next *ListNode
}

func createList(slice []int32) *ListNode {
	var list []*ListNode

	for index, value := range slice {
		list = append(list, new(ListNode))
		list[index].Val = value
		if index > 0 {
			list[index-1].Next = list[index]
		}
	}

	return list[0]
}

func deleteDuplicates(list *ListNode) {
	var cur *ListNode = list

	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
}

func main() {
	var list ListNode = *createList([]int32{1, 2, 3, 3, 3, 3, 4, 4})

	var cur *ListNode = &list
	for cur != nil {
		fmt.Printf("%v\n", cur.Val)
		cur = cur.Next
	}

	deleteDuplicates(&list)

	fmt.Printf("After deletion:\n")
	cur = &list
	for cur != nil {
		fmt.Printf("%v\n", cur.Val)
		cur = cur.Next
	}
}
