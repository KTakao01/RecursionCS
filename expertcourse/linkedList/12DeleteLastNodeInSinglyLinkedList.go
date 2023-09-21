package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func deleteTail(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	// 関数を完成させてください
	if head.next == nil {
		head = nil
		return nil
	}
	iterator := head
	for iterator.next.next != nil {
		iterator = iterator.next
	}
	iterator.next = nil
	return head
}
