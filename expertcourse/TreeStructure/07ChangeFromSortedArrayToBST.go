package main

type Integer struct {
	value int32
}

type BinaryTree struct {
	data  *Integer
	left  *BinaryTree
	right *BinaryTree
}

func sortedArrToBST(numberList []int32) *BinaryTree {
	// 関数を完成させてください
	if len(numberList) == 0 {
		return nil
	}

	return sortedArrToBSTHelper(numberList, 0, len(numberList)-1)
}

func sortedArrToBSTHelper(numberList []int32, start, end int) *BinaryTree {
	if start == end {
		return &BinaryTree{data: &Integer{value: numberList[start]}, left: nil, right: nil}
	}

	mid := (start + end) / 2
	var left, right *BinaryTree
	if start <= mid-1 {
		left = sortedArrToBSTHelper(numberList, start, mid-1)
	}
	if end >= mid+1 {
		right = sortedArrToBSTHelper(numberList, mid+1, end)
	}

	root := &BinaryTree{data: &Integer{value: numberList[mid]}, left: left, right: right}
	return root

}
