/*
https://recursionist.io/dashboard/course/3/lesson/492
整数によって構成される intArr が与えられるので、ヒープソートアルゴリズムによって、昇順ソートする、heapsort という関数を作成します。配列の要素間の入れ替えをすることによって、空間計算量 O(1) で実装してください。
*/

package main

import (
	"math"
)

func heapsort(intArr []int32) []int32 {
	// 関数を完成させてください
	//配列全体を最大ヒープにする
	buildMaxHeap(intArr)
	heapEnd := int32(len(intArr) - 1)
	for heapEnd > 0 {
		//最大ヒープの最大値を最後の葉ノードとスワップする
		intArr[0], intArr[heapEnd] = intArr[heapEnd], intArr[0]
		//入力用スライスのheapEndを-1する
		heapEnd--
		//根ノードに最大ヒープにする関数を適用する
		maxHeapify(intArr, heapEnd, 0)
	}
	return intArr
}

func maxHeapify(intArr []int32, heapEnd int32, i int32) {
	//対象ノード
	biggest := i
	//左の子ノード
	l := left(i)
	if l <= heapEnd && intArr[l] > intArr[biggest] {
		biggest = l
	}
	//右の子ノード
	r := right(i)
	if r <= heapEnd && intArr[r] > intArr[biggest] {
		biggest = r
	}

	//スワップ
	if biggest != i {
		intArr[i], intArr[biggest] = intArr[biggest], intArr[i]
		//スワップ後、最大ヒープ化
		maxHeapify(intArr, heapEnd, biggest)
	}
}

func buildMaxHeap(intArr []int32) {
	middle := parent(int32(len(intArr)))
	for i := middle; i >= 0; i-- {
		maxHeapify(intArr, int32(len(intArr)-1), i)
	}
}

func parent(i int32) int32 {
	return int32(math.Floor((float64(i-1) / 2)))
}

func left(i int32) int32 {
	return 2*i + 1
}

func right(i int32) int32 {
	return 2*i + 2
}
