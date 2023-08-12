package linkedList

import (
	"fmt"
)

type Item struct {
	data int
	next *Item
}

func NewItem(data int) *Item {
	return &Item{data: data}
}

type SinglyLinkedList struct {
	head *Item
}

func NewSinglyLinkedList(item *Item) *SinglyLinkedList {
	return &SinglyLinkedList{head: item}
}

func main() {
	item1 := NewItem(7)
	item2 := NewItem(99)
	item3 := NewItem(45)

	numList := NewSinglyLinkedList(item1)

	item1.next = item2
	item2.next = item3

	currentNode := numList.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}
