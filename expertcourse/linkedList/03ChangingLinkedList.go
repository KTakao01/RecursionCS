package main

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func NewSinglyLinkedListNode(data int32) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{data: data}
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
}

func NewSinglyLinkedList(data int32) *SinglyLinkedList {
	return &SinglyLinkedList{head: &SinglyLinkedListNode{data: data}}
}

func getLinkedList(arr []int32) *SinglyLinkedListNode {
	//初期値
	numList := NewSinglyLinkedList(arr[0])
	currentNode := numList.head

	//ノード作成
	for _, val := range arr[1:] {
		currentNode.next = NewSinglyLinkedListNode(val)
		currentNode = currentNode.next
	}

	//currentNodeがループにより最後のリストを示すので再度リセット
	currentNode = numList.head
	var output string = fmt.Sprintf("%d", currentNode.data)
	//出力
	for currentNode != nil {
		if currentNode.next != nil {
			output += fmt.Sprintf("➡%d", currentNode.next.data)
		}
		currentNode = currentNode.next
	}

	output += "➡END"
	fmt.Println(output)
	return numList.head

}
