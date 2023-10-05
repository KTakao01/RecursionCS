package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func bstSearch(root *BinaryTree, key int32) *BinaryTree {
	// 関数を完成させてください
	if root == nil {
		return nil
	}

	if root.data.value == key {
		return root
	}

	if root.data.value > key {
		return bstSearch(root.left, key)
	}

	if root.data.value < key {
		return bstSearch(root.right, key)
	}

	return nil
}
