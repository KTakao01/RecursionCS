//https://recursionist.io/dashboard/course/3/lesson/460

package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func minimumNode(root *BinaryTree) *BinaryTree {
	// 関数を完成させてください
	iterator := root
	var min int32 = root.data.value
	minimumNode := iterator

	for iterator != nil {
		if iterator.data.value < min {
			min = iterator.data.value
			minimumNode = iterator
		}
		iterator = iterator.left
	}
	return minimumNode
}
