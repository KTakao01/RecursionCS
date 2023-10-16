// https://recursionist.io/dashboard/problems/345
// 二分木 root が与えられるので、各階層の一番右側にあるノードを返す、rightLevelNode という関数を作成してください。
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

func rightLevelNode(root *BinaryTree) []int32 {
	// 関数を完成させてください
	if root == nil {
		return nil
	}
	queue := []*BinaryTree{}
	queue = append(queue, root)
	results := []int32{}
	results = append(results, root.data.value)
	return rightLevelNodeHelper(root, queue, results)
}

func rightLevelNodeHelper(root *BinaryTree, queue []*BinaryTree, results []int32) []int32 {
	fmt.Println("初期results", results)
	child := []*BinaryTree{}

	for _, node := range queue {
		if node.left != nil {
			fmt.Println("node.left", node.left.data.value)
			child = append(child, node.left)
		}

		if node.right != nil {
			fmt.Println("node.right", node.right.data.value)
			child = append(child, node.right)
		}
	}

	fmt.Println("child", child)
	if len(child) != 0 {
		fmt.Println("childあり", child[len(child)-1].data.value)
		results = append(results, child[len(child)-1].data.value)
		return rightLevelNodeHelper(root, child, results)
	}
	return results
}

//幅階層走査
//左→右と走査するな
//queueの一番最後尾がその階層の最右端では？
