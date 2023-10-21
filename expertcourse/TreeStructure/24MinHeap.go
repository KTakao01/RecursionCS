/*
https://recursionist.io/dashboard/course/3/lesson/491
整数によって構成される intArr が与えられるので、最小値が先頭に配置される最小ヒープを表す配列を返す、buildMinHeap という関数を作成してください。
*/

package main

import (
	"fmt"
	"math"
)

// 特定のノードとその子ノードだけを最小ヒープにする
func minHeapify(arr []int32, heapEnd int32, i int32) {
	//親ノード
	smallest := i
	//左の子ノード
	left := 2*smallest + 1
	right := 2*smallest + 2
	if left <= heapEnd && arr[smallest] > arr[left] {
		smallest = left
	}
	//右の子ノード
	if right <= heapEnd && arr[smallest] > arr[right] {
		smallest = right
	}
	//最小値のインデックスbiggestと対象のインデックスiをスワップ。最大値のインデックスiと対象のインデックスsmallestになる。
	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		minHeapify(arr, heapEnd, smallest)
	}
}

// 配列全体を最小ヒープにする
func buildMinHeap(intArr []int32) []int32 {
	// 関数を完成させてください
	i := len(intArr) - 1
	middle := int32(math.Floor(float64((i - 1) / 2)))
	heapEnd := int32(len(intArr) - 1)
	heapArr := make([]int32, len(intArr))
	copy(heapArr, intArr)
	//葉ノード以外をminHeapifyで最小ヒープ
	for i := middle; i >= 0; i-- {
		minHeapify(heapArr, heapEnd, i)
		fmt.Println(heapArr)
	}
	return heapArr
}
