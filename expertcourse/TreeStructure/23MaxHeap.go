// https://recursionist.io/dashboard/course/3/lesson/490
// 整数によって構成される intArr が与えられるので、最大値が先頭に配置される最大ヒープを表す配列を返す、buildMaxHeap という関数を作成してください
package main

import (
	"fmt"
	"math"
)

// 特定のノードとその子ノードだけを最大ヒープにする
func maxHeapify(arr []int32, heapEnd int32, i int32) {
	//親ノード
	biggest := i
	//左の子ノード
	left := 2*biggest + 1
	right := 2*biggest + 2
	if left <= heapEnd && arr[biggest] < arr[left] {
		biggest = left
	}
	//右の子ノード
	if right <= heapEnd && arr[biggest] < arr[right] {
		biggest = right
	}
	//最大値のインデックスbiggestと対象のインデックスiをスワップ。最大値のインデックスiと対象のインデックスbiggestになる。
	if biggest != i {
		arr[i], arr[biggest] = arr[biggest], arr[i]
		maxHeapify(arr, heapEnd, biggest)
	}
}

// 配列全体を最大ヒープにする
func buildMaxHeap(intArr []int32) []int32 {
	// 関数を完成させてください
	i := len(intArr) - 1
	middle := int32(math.Floor(float64((i - 1) / 2)))
	heapEnd := int32(len(intArr) - 1)
	heapArr := make([]int32, len(intArr))
	copy(heapArr, intArr)
	//葉ノード以外をmaxHeapifyで最大ヒープ
	for i := middle; i >= 0; i-- {
		maxHeapify(heapArr, heapEnd, i)
		fmt.Println(heapArr)
	}
	return heapArr
}
