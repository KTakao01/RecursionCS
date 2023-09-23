package main

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func findMergeNode(headA *SinglyLinkedListNode, headB *SinglyLinkedListNode) int32 {
	// 関数を完成させてください
	headA = reverse(headA)
	headB = reverse(headB)
	fmt.Println(headA, headB)
	iteratorA := headA
	iteratorB := headB
	//AとBの交わりが存在しない場合
	if iteratorA.data != iteratorB.data {
		return -1
	}
	//AとBの交わりが存在する場合
	for iteratorA != nil || iteratorB != nil {
		preA := iteratorA
		iteratorA = iteratorA.next
		iteratorB = iteratorB.next

		if iteratorA.data != iteratorB.data {
			return preA.data
		}
	}
	return -9999
}

func reverse(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	current := head
	var pre *SinglyLinkedListNode

	for current != nil {
		nextNode := current.next
		current.next = pre
		pre = current
		current = nextNode
		fmt.Println(pre, current)
	}
	head = pre
	return head
}

//考え方
//前方から追跡していくのは同じステップ刻むのでサイズの異なるリストに対しては無効
//後ろから辿る。後ろからたどり、最初に一致しない場所のnextが求めるノードの値
//上記は双方向リストでしか無理
//リストを逆順にする
//走査してそれぞれのノードの値を順に比較
//異なる値がでた初めてのノードの前の値を返す
