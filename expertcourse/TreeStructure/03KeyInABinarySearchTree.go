//https://recursionist.io/dashboard/course/3/lesson/459

package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func exists(root *BinaryTree, key int32) bool {
	// 関数を完成させてください
	if root == nil || root.data == nil {
		return false
	}

	if root.data.value > key {
		return exists(root.left, key)
	}

	if root.data.value < key {
		return exists(root.right, key)
	}

	return root.data.value == key
}
