package main

import (
	"fmt"
)

type Item struct {
	data int
	prev *Item
	next *Item
}

// ノードを作成する
func NewItem(data int) *Item {
	return &Item{data: data}
}

type DoublyLinkedList struct {
	head *Item
	tail *Item
}

// 双方向リストを構築する
func NewDoublyLinkedList(arr []int) *DoublyLinkedList {
	if len(arr) == 0 {
		return &DoublyLinkedList{}
	}
	list := &DoublyLinkedList{head: NewItem(arr[0])}
	currentNode := list.head
	currentNode.prev = nil
	for i := 1; i < len(arr); i++ {
		currentNode.next = NewItem(arr[i])
		currentNode.next.prev = currentNode
		currentNode = currentNode.next
	}
	list.tail = currentNode
	return list
}

func main() {
	numList := NewDoublyLinkedList([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(numList.head.data)
	fmt.Println(numList.head.next.data)
	fmt.Println(numList.head.next.prev.data)
	fmt.Println(numList.tail.data)
	fmt.Println(numList.tail.prev.data)
	fmt.Println(numList.tail.prev.prev.data)
}
