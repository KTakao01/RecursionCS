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

// delete,findParent,search,transplant,successorの機能が必要
func bstDelete(root *BinaryTree, key int32) *BinaryTree {
	// 関数を完成させてください
	if root == nil {
		return nil
	}
	node := search(root, key)
	parent := findParent(root, node)
	if parent == nil && root != node {
		// handle the error appropriately. For example:
		fmt.Println("Error: Parent not found and node is not root")
		return root
	}
	var newRoot *BinaryTree
	if node == nil {
		return root
	}

	//case1 ノードNが葉ノード

	if node.left == nil && node.right == nil {
		if parent.left == node {
			parent.left = nil
		} else if parent.right == node {
			parent.right = nil
		}
	} else if node.left == nil {
		fmt.Println("a")
		newRoot = transPlant(root, parent, node, node.right)
		if newRoot != root {
			return newRoot
		}
	} else if node.right == nil {
		newRoot = transPlant(root, parent, node, node.left)
		if newRoot != root {
			return newRoot
		}
	} else {
		successorNode := findSuccessor(root, node)
		fmt.Println(successorNode.data.value)
		successorParent := findParent(root, successorNode)
		if node.right != successorNode {

			newRoot = transPlant(root, successorParent, successorNode, successorNode.right)
			//ノードN、移植先の後続ノードNの右部分木の更新
			successorNode.right = node.right
		}
		newRoot = transPlant(root, parent, node, successorNode)
		successorNode.left = node.left
		return newRoot
	}
	return root
}

// nodeParentの子ノードであるnodeをtargetに置き換える。
func transPlant(root, nodeParent, node, target *BinaryTree) *BinaryTree {
	//以下は盲点だった。親ノードがnilのとき子ノードは根ノード。子ノードにtargetを移植する
	if nodeParent == nil {
		fmt.Println("b")
		return target
	} else if nodeParent.left == node {
		nodeParent.left = target
	} else {
		nodeParent.right = target
	}
	return root

}

// 後続ノードを返す関数
func findSuccessor(root, node *BinaryTree) *BinaryTree {
	//ノードに右部分木がある場合、最左端
	fmt.Println("target", node.data.value)
	if node.right != nil {
		fmt.Println("c")
		return minimumNode(node.right)
	}
	//ノードに右部分木がない場合、祖先ノード
	iterator := root
	var successor *BinaryTree
	for iterator != nil && node != iterator {
		if iterator.data.value > node.data.value {
			successor = iterator
			iterator = iterator.left
		} else {
			iterator = iterator.right
		}
	}
	//if iterator == node {
	//  return nil
	//}
	return successor
}

func minimumNode(root *BinaryTree) *BinaryTree {
	iterator := root
	//if iterator == nil || iterator.left == nil{
	//    return iterator
	//}
	for iterator != nil && iterator.left != nil {
		iterator = iterator.left
	}
	return iterator
}

// BST内にkeyがあるときのその部分木を返す関数
func search(root *BinaryTree, key int32) *BinaryTree {
	iterator := root
	for iterator != nil {
		if iterator.data.value == key {
			return iterator
		} else if iterator.data.value > key {
			iterator = iterator.left
		} else {
			iterator = iterator.right
		}
	}
	return nil
}

// ノードの親ノードを返す関数
func findParent(root, node *BinaryTree) *BinaryTree {
	if node == nil || root == node {
		return nil // ノードがnilまたはrootである場合、親は存在しません
	}
	iterator := root
	for iterator != nil {
		if iterator.left != nil && iterator.left.data != nil && iterator.left.data.value == node.data.value {
			return iterator
		} else if iterator.right != nil && iterator.right.data != nil && iterator.right.data.value == node.data.value {
			return iterator
		} else if node.data != nil {
			if iterator.data.value > node.data.value {
				iterator = iterator.left
			} else if iterator.data.value < node.data.value {
				iterator = iterator.right
			}
		}
	}
	return nil
}
