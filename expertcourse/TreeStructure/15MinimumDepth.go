// https://recursionist.io/dashboard/problems/277
package main

import (
	"fmt"
)

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func minDepth(root *BinaryTree) int32 {
	// 関数を完成させてください
	var queue []*BinaryTree
	if root != nil {
		queue = append(queue, root)
	}
	return minDepthHelper(root, queue, 1)

}

func minDepthHelper(root *BinaryTree, queue []*BinaryTree, count int32) int32 {
	child := []*BinaryTree{}
	fmt.Println("queue", len(queue))
	fmt.Println("count-pre", count)
	for _, node := range queue {
		if (node.left == nil || node.left.data == nil) && (node.right == nil || node.right.data == nil) {
			fmt.Println("a")
			return count
		}
		if node.left != nil {
			child = append(child, node.left)
		}
		if node.right != nil {
			child = append(child, node.right)
		}
	}

	return minDepthHelper(root, child, count+1)

}

//葉ノードまでの距離
//最深層までの葉ノードの合計の解き方を応用する
//
