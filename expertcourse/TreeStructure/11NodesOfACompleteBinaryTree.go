// https://recursionist.io/dashboard/problems/304
package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func countNodes(root *BinaryTree) int32 {
	// 関数を完成させてください
	if root == nil {
		return 0
	}
	return countNodesHelper(root, 1)
}

func countNodesHelper(root *BinaryTree, count int32) int32 {
	if root != nil {
		if root.left != nil {
			count = countNodesHelper(root.left, count+1)
		}
		if root.right != nil {
			count = countNodesHelper(root.right, count+1)
		}
	}
	return count
}
