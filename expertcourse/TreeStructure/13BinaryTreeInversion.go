// https://recursionist.io/dashboard/problems/submissions/680147
package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func invertTree(root *BinaryTree) *BinaryTree {
	// 関数を完成させてください
	if root == nil || root.data == nil {
		return nil
	}
	reverseR := &Integer{root.data.value}
	rRight := invertTree(root.left)
	rLeft := invertTree(root.right)
	return &BinaryTree{reverseR, rLeft, rRight}
}
