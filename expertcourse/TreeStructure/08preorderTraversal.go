// https://recursionist.io/dashboard/course/3/lesson/469
package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func preorderTraversal(root *BinaryTree) []int32 {
	// 関数を完成させてください
	results := []int32{}
	results = preorderWalk(root, results)
	return results
}

func preorderWalk(root *BinaryTree, results []int32) []int32 {
	if root == nil || root.data == nil {
		return results
	}
	results = append(results, root.data.value)
	results = preorderWalk(root.left, results)
	results = preorderWalk(root.right, results)
	return results
}
