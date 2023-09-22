package main

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertAtPosition(head *SinglyLinkedListNode, position int32, data int32) *SinglyLinkedListNode {
	// 関数を完成させてください
	dummyNode := &SinglyLinkedListNode{data: 0}
	dummyNode.next = head
	iterator := head
	var count int32
	for i := int32(0); i < position && iterator != nil; i++ {
		iterator = iterator.next
		count++
		fmt.Println("count", count)
	}

	if iterator != nil {
		temp := iterator.next
		iterator.next = &SinglyLinkedListNode{data: data}
		iterator.next.next = temp
	}

	return head

}
