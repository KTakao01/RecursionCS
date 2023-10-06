// https://recursionist.io/dashboard/course/3/lesson/460
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
	if root.left == nil {
		return root
	}

	return minimumNode(root.left)
}
