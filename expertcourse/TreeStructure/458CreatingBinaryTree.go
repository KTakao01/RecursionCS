package main

import (
	"fmt"
)

type binaryTree struct {
	data  int
	left  *binaryTree
	right *binaryTree
}

func NewBinaryTree(data int) *binaryTree {
	return &binaryTree{data: data}
}

func main() {
	binaryTree := NewBinaryTree(10)
	node2 := NewBinaryTree(6)
	node3 := NewBinaryTree(8)

	binaryTree.left = node2
	binaryTree.right = node3

	fmt.Println(binaryTree.data)
	fmt.Println(binaryTree.left.data)
	fmt.Println(binaryTree.right.data)
}
