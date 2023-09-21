package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func linkedListLastValue(head *SinglyLinkedListNode) int32 {
	// 関数を完成させてください
	iterator := head
	for iterator.next != nil {
		iterator = iterator.next
	}
	return iterator.data
}
