package linkedList

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

type SinglyLinkedList struct {
	head *Node
}

func NewSinglyLinkedList(data int) *SinglyLinkedList {
	return &SinglyLinkedList{head: &Node{data: data}}
}

func main() {
	array := []int{35, 23, 546, 67, 86, 234, 56, 767, 34, 1, 98, 78, 555}

	//for i:=0,i<array.length(),i++{

	//}

	numList := NewSinglyLinkedList(array[0])
	currentNode := numList.head

	//リストに他のノードを追加します
	for _, val := range array[1:] {
		currentNode.next = NewNode(val)
		currentNode = currentNode.next
	}

	currentNode = numList.head
	//連結リストを反復して出力する
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}
