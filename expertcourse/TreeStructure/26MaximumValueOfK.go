/*
https://recursionist.io/dashboard/course/3/lesson/493
整数によって構成される intArr と整数 k（0 < k <= intArr.length）が与えられるので、配列から k 個の最大値を返す、topKElements という関数を作成してください。
*/

package main

import (
	"math"
)

func topKElements(intArr []int32, k int32) []int32 {
	// 関数を完成させてください
	heapsort(intArr)
	results := []int32{}
	heapEnd := int32(len(intArr) - 1)
	for i := int32(0); i < k; i++ {
		results = append(results, intArr[heapEnd-i])
	}
	return results
}

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
