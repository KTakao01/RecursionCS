package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
}

func NewSinglyLinkedList(arr []int) *SinglyLinkedList {
	if len(arr) == 0 {
		return &SinglyLinkedList{head: nil}
	}

	head := &Node{data: arr[0]}
	currentNode := head

	for i := 1; i < len(arr); i++ {
		currentNode.next = &Node{data: arr[i]}
		currentNode = currentNode.next
	}

	return &SinglyLinkedList{head: head}
}

func (s *SinglyLinkedList) At(index int) *Node {
	iterator := s.head
	for i := 0; i < index; i++ {
		if iterator == nil {
			return nil
		}
		iterator = iterator.next
	}
	return iterator
}

// TODO: Implement the findNode function here.
func (s *SinglyLinkedList) findNode(key int) int {
	firstNum := s.head
	var count int = 0
	for firstNum.next != nil {
		if firstNum.data == key {
			return count
		}

		//if firstNum,ok = firstNum.next;!ok
		if firstNum.next != nil {
			firstNum = firstNum.next
			count++
		} else {
			break
		}

	}
	return -1
}

func main() {
	arr := []int{35, 23, 546, 67, 86, 234, 56, 767, 34, 1, 98, 78, 555}
	numList := NewSinglyLinkedList(arr)

	// Example usage of the At function
	nodeAt2 := numList.At(2)
	if nodeAt2 != nil {
		fmt.Println(nodeAt2.data) // Should print 546
	}
	fmt.Println(numList.findNode(767))
}
