// https://recursionist.io/dashboard/course/3/lesson/479
// 二分木 root が与えられるので、それが左右対称かどうかを確かめる、symmetricTree という関数を作成してください。
package main

import (
	"fmt"
	"reflect"
)

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func symmetricTree(root *BinaryTree) bool {
	// 関数を完成させてください
	var results []int32 = []int32{}
	var leftArr, rightArr []int32 = []int32{}, []int32{}
	if root == nil {
		return true
	}

	if root.left != nil {
		leftArr = leftSubTree(root.left, results)
	}
	fmt.Println("leftArr", leftArr)
	if root.right != nil {
		rightArr = rightSubTree(root.right, results)
	}
	fmt.Println("rightArr", rightArr)

	return reflect.DeepEqual(leftArr, rightArr)
}

// 左右ループ
func leftSubTree(root *BinaryTree, results []int32) []int32 {
	if root == nil {
		results = append(results, -999)
		return results
	} else {
		results = append(results, root.data.value)
	}

	fmt.Println("a", results)
	results = leftSubTree(root.left, results)
	fmt.Println("a-2", results)

	fmt.Println("b", results)
	results = leftSubTree(root.right, results)
	fmt.Println("b-2", results)

	return results
}

func rightSubTree(root *BinaryTree, results []int32) []int32 {
	if root == nil {
		results = append(results, -999)
		return results
	} else {
		results = append(results, root.data.value)
	}

	fmt.Println("b", results)
	results = rightSubTree(root.right, results)
	fmt.Println("b-2", results)

	fmt.Println("a", results)
	results = rightSubTree(root.left, results)
	fmt.Println("a-2", results)

	return results
}
