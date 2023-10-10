// https://recursionist.io/dashboard/course/3/lesson/470
package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func inorderTraversal(root *BinaryTree) []int32 {
	// 関数を完成させてください
	results := []int32{}
	return inorderWalk(root, results)
}

func inorderWalk(root *BinaryTree, results []int32) []int32 {
	if root != nil && root.data != nil {
		results = inorderWalk(root.left, results)
		results = append(results, root.data.value)
		results = inorderWalk(root.right, results)
	}
	return results
}
