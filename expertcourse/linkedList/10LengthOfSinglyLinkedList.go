package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func linkedListLength(head *SinglyLinkedListNode) int32 {
	// 関数を完成させてください
	iterator := head
	var count int32 = 0
	for iterator != nil {
		iterator = iterator.next
		count++
	}
	return count
}
