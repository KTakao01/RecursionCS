// https://recursionist.io/dashboard/problems/275
package main

import (
	"math"
)

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func allElementsSorted(root1 *BinaryTree, root2 *BinaryTree) []int32 {
	// 関数を完成させてください
	var results1, results2, merged []int32
	results1 = inorderTraverse(root1, results1)
	results2 = inorderTraverse(root2, results2)
	merged = make([]int32, 0, len(results1)+len(results2))
	results1 = append(results1, math.MaxInt32)
	results2 = append(results2, math.MaxInt32)
	var i, j int = 0, 0

	for i+j < len(results1)+len(results2)-2 {
		if results1[i] <= results2[j] {
			merged = append(merged, results1[i])
			i++
		} else {
			merged = append(merged, results2[j])
			j++
		}
	}

	return merged
}

func inorderTraverse(root *BinaryTree, results []int32) []int32 {
	if root == nil || root.data == nil {
		return results
	}
	results = inorderTraverse(root.left, results)
	results = append(results, root.data.value)
	results = inorderTraverse(root.right, results)
	return results
}

//スタックやキューを用いて一方のスライスのループで解こうとしたが格納できる数値が少ないので制御できないパターンが出る。
//両方のスライスを同時に扱う方法を考える。
