package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func NewSinglyLinkedListNode(data int32) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{data: data}
}

func insertHeadTail(head *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	// 関数を完成させてください
	newHead := NewSinglyLinkedListNode(data)
	newTail := NewSinglyLinkedListNode(data)
	newHead.next = head
	iterator := newHead
	for iterator.next != nil {
		iterator = iterator.next
	}
	iterator.next = newTail
	return newHead

}
