//https://grecursionist.io/dashboard/problems/278
/*二分木が与えられるので、完全二分木（complete binary tree）かどうか判定する、isCompleteBinaryTree という関数を作成してください。完全二分木とは、最下層を除いてすべての深さがノードで満たされ、最下層の葉ノードが可能な限り左に寄せられているような木のことを指します。
 */

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

func isCompleteBinaryTree(root *BinaryTree) bool {
	// 関数を完成させてください
	if root == nil {
		return false
	}
	queue := []*BinaryTree{root}
	nilCheck := true
	var flag int32 = 0
	return isCompleteBinaryTreeHelper(root, queue, nilCheck, flag)
}

func isCompleteBinaryTreeHelper(root *BinaryTree, queue []*BinaryTree, nilCheck bool, flag int32) bool {
	child := []*BinaryTree{}
	if nilCheck == false {
		return false
	}
	for _, node := range queue {
		if node.left != nil {
			child = append(child, node.left)
		} else {
			flag = 1
		}

		if flag == 1 && node.left != nil {
			child = append(child, node.left)
			fmt.Println("a")
			nilCheck = false
		}

		if node.right != nil {
			child = append(child, node.right)
		} else {
			flag = 1
		}

		if flag == 1 && node.right != nil {
			child = append(child, node.right)
			fmt.Println("b")
			nilCheck = false
		}
	}

	if len(child) > 0 && nilCheck == true {
		return isCompleteBinaryTreeHelper(root, child, nilCheck, flag)
	}
	return nilCheck
}

//幅階層走査
//左→右に走査

//最下層の葉ノードがなるべく左に寄せられている木という条件
//nilもresultにいれられたら。。。配列を確認すれば良い。
//配列 *BinaryTree
//for rangeでnilないか確認
//あればfalse なければtrue

//queueを確認

//最下層を除いて全ての深さがノードでみたされる。
//childを確認して、nil？
//nilにならない限りchildは左右に子ノードがあるから偶数になる。
//もしくはnilになった場合の条件分岐をつくりフラグを設定する。

//終了条件のnilチェックが変更される
//全てのchildがnilのとき終了？
