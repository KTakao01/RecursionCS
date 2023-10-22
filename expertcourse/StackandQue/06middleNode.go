package main

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func middleNode(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	// 関数を完成させてください
	iterator := head
	var count int = 1
	for iterator.next != nil {
		iterator = iterator.next
		count++
		fmt.Println("count:", count)
	}
	iterator = head
	var i int = 1
	for iterator.next != nil {
		if i == count/2+1 {
			head = iterator
			fmt.Println(i)
		}
		fmt.Println("処理済み")
		iterator = iterator.next
		i++
	}
	if count == 2 {
		return head.next
	}
	return head
}
