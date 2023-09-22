package main

import (
	"math"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertNodeInSorted(head *SinglyLinkedListNode, data int32) *SinglyLinkedListNode {
	// ダミーのノードを作ります。dataは仮にintの最小値MinInt32を入れておきます。
	dummyNode := &SinglyLinkedListNode{data: math.MinInt32}
	// ダミーノードをhead の前に挿入します。
	dummyNode.next = head

	iterator := dummyNode

	// 挿入すべき位置までリストを走査します。
	for iterator.next != nil && iterator.next.data < data {
		iterator = iterator.next
	}

	// 新しいノードを作ります。
	node := &SinglyLinkedListNode{data: data}
	// iterator.nextをtempに入れ保管します。
	temp := iterator.next
	// iterator.nextには新しいノードを入れます。
	iterator.next = node
	// tempに入れていた参照を新しいノードのnextにします。
	node.next = temp

	// ダミーの次のノードから返します。
	return dummyNode.next
}
