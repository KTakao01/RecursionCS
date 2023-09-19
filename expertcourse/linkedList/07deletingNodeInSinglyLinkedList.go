package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func NewSinglyLinkedListNode(data int32) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{data: data}
}

func removeNthNode(head *SinglyLinkedListNode, n int32) *SinglyLinkedListNode {
	iterator := head
	count := int32(1)

	// headを含む連結リストの長さを求める
	for iterator.next != nil {
		iterator = iterator.next
		count++
	}

	// nがリストの長さ以上の場合、またはnが0の場合、headをそのまま返す
	if count < n || n == 0 {
		return head
	}

	// nがリストの長さと同じ場合、headを1つ進めて返す
	if count == n {
		return head.next
	}

	iterator = head
	if n == 1 {
		for i := int32(0); i < count-n-1; i++ {
			iterator = iterator.next
		}
		iterator.next = nil
	} else {
		for i := int32(0); i < count-n-1; i++ {
			iterator = iterator.next
		}
		iterator.next = iterator.next.next
	}

	return head
}
