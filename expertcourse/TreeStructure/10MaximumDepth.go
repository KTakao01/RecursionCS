// https://recursionist.io/dashboard/course/3/lesson/473
package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func maximumDepth(root *BinaryTree) int32 {
	// 関数を完成させてください
	if root == nil || root.data == nil {
		return 0
	}

	leftDepth := maximumDepth(root.left)

	rightDepth := maximumDepth(root.right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

/*
・−３のノードを探索するとき左も右もroot.data　== nilなので0を返す。
・leftDepth < rightDepthよりrightDepth + 1を返すことから1を返す
・ここで-10のノードを探索に戻るが、rightDepth  = 1 を引き継ぎ、左の子ノードはnilなのでleftDepth=1
・・leftDepth <= rightDepthよりrightDepth + 1を返すことから２を返す。これをleftDepthが保存する。
・右部分木においても同様に探索して-10のノードには2を返す。これをrightDepthが保存する
・leftDepth <= rightDepthよりrightDepth + 1を返すことから3を返す。
*/
