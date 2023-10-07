// https://recursionist.io/dashboard/course/3/lesson/462
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

func successor(root *BinaryTree, key int32) *BinaryTree {
	// 関数を完成させてください
	if root == nil {
		fmt.Println("nil")
		return nil
	}

	// 右部分木がない場合があるので後続ノードは全体から探す
	results := []*Integer{}
	bigger(root, key, &results)
	iterator := root
	for len(results) != 0 && iterator != nil {
		if iterator.data.value == results[0].value {
			return iterator
		}

		if iterator.data.value > results[0].value {
			iterator = iterator.left
		} else {
			iterator = iterator.right
		}
	}
	return nil
}

// キーの右部分木がないってだけでキーのノード自体は右部分木に位置している場合もあるため全体から探索するほかなさそう
// 後続ノードの値を全探索して返す
func bigger(root *BinaryTree, key int32, results *[]*Integer) {
	fmt.Println("bigger", key)
	if root == nil || root.data == nil {
		fmt.Println("a", root)
		return
	}
	//探索は再帰
	if root != nil || root.data != nil {
		bigger(root.left, key, results)

		fmt.Println("b", root.data.value)
		if root.data.value > key {
			fmt.Println("追加", root.data.value)
			*results = append(*results, root.data)
		}

		bigger(root.right, key, results)
	}

}
