package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func reverseLinkedList(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	// 関数を完成させてください
	var prev *SinglyLinkedListNode = nil
	current := head

	for current != nil {
		//nextNodは次に処理するノードを保存します。
		nextNod := current.next
		//currentノードの次のノードを保存。このとき逆順に更新される。
		current.next = prev
		//current,prevを次のリストに進める。
		prev = current
		current = nextNod
	}
	return prev
}
