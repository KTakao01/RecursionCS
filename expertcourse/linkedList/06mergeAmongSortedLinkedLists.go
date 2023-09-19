package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func NewSinglyLinkedListNode(data int32) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{data: data}
}

func mergeTwoSortedLinkedLists(head1 *SinglyLinkedListNode, head2 *SinglyLinkedListNode) *SinglyLinkedListNode {
	// 関数を完成させてください
	var newHead *SinglyLinkedListNode
	if head1.data <= head2.data {
		newHead = NewSinglyLinkedListNode(head1.data)
		head1 = head1.next
	} else {
		newHead = NewSinglyLinkedListNode(head2.data)
		head2 = head2.next
	}

	iterator3 := newHead

	for head1 != nil || head2 != nil {
		if head1 == nil {
			iterator3.next = head2
			head2 = head2.next

		} else if head2 == nil {
			iterator3.next = head1
			head1 = head1.next
		} else {
			if head1.data >= head2.data {
				iterator3.next = head2
				head2 = head2.next
			} else {
				iterator3.next = head1
				head1 = head1.next
			}

		}
		iterator3 = iterator3.next
	}
	return newHead
}
