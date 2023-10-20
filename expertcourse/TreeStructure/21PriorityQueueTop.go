//https://recursionist.io/dashboard/problems/451
/*
優先度付きキューは特定の操作に基づくデータ構造で、それらの操作は以下の通りです。

top: 最も高い優先度を持つ要素を取得する
pop: 最も高い優先度を持つ要素を削除する（今回の課題では実装が必要ない）
insert: 要素を挿入する（今回の課題では実装が必要ない）
これらの操作を持つ優先度付きキューは、ソート済みの配列を使っても実装できますが、最も高い優先度の要素に焦点を当てるデータ構造のため、より効率的なヒープが使用されます。

整数の配列 arr と配列から最大ヒープを構築する HeapLibrary が与えられたとき、以下の条件に従って PriorityQueue クラスを作成してください。

ArrayList maxHeap
動的配列で、優先度付きキューの要素を保持します。この配列は最大ヒープと呼ばれる特殊な形状に整列します。これにより、最大の要素が常に根に存在することが保証されます。arr が与えられたとき、その状態を変更せずにこの配列を最大ヒープに変換する必要があります。arr のディープコピーを作成し元の配列の変更を防ぎつつ、HeapLibrary の buildMaxHeap 関数を使用してそのコピーを最大ヒープに変換することで達成できます。

int top()
最大ヒープの最大値を返します。
*/
package main

import (
	"fmt"
	"math"
)

// HeapLibrary contains utility functions for working with a heap.
type HeapLibrary struct{}

// BuildMaxHeap constructs a max heap from the provided slice.
func (hl *HeapLibrary) BuildMaxHeap(arr []int) {
	middle := hl.Parent(len(arr) - 1)
	for i := middle; i >= 0; i-- {
		hl.MaxHeapify(arr, len(arr)-1, i)
	}
}

// MaxHeapify ensures the heap property is maintained at the specified node and its descendants.
func (hl *HeapLibrary) MaxHeapify(arr []int, heapEnd, i int) {
	l := hl.Left(i)
	r := hl.Right(i)

	biggest := i
	if l <= heapEnd && arr[l] > arr[biggest] {
		biggest = l
	}
	if r <= heapEnd && arr[r] > arr[biggest] {
		biggest = r
	}

	if biggest != i {
		arr[i], arr[biggest] = arr[biggest], arr[i]
		hl.MaxHeapify(arr, heapEnd, biggest)
	}
}

// Left returns the index of the left child of the specified node.
func (hl *HeapLibrary) Left(i int) int {
	return 2*i + 1
}

// Right returns the index of the right child of the specified node.
func (hl *HeapLibrary) Right(i int) int {
	return 2*i + 2
}

// Parent returns the index of the parent of the specified node.
func (hl *HeapLibrary) Parent(i int) int {
	return int(math.Floor(float64(i-1) / 2))
}

type PriorityQueue struct {
	maxHeap []int
}

func NewPriorityQueue(arr []int) *PriorityQueue {
	maxHeap := make([]int, len(arr))
	copy(maxHeap, arr)
	heapLib := HeapLibrary{}
	heapLib.BuildMaxHeap(maxHeap)
	return &PriorityQueue{maxHeap: maxHeap}

}

func (pq *PriorityQueue) top() int {
	fmt.Println(pq.maxHeap[0])
	return pq.maxHeap[0]
}
func main() {
	arr1 := []int{2, 3, 43, 2, 53, 6, 75, 10}
	pq1 := NewPriorityQueue(arr1)
	pq1.top()

	arr2 := []int{3, 12, 0, 2, 9, 1, 65, 32}
	pq2 := NewPriorityQueue(arr2)
	pq2.top()

	arr3 := []int{1, 2, 3, 4, 8, 2, 1, 9, 7, 3, 4}
	pq3 := NewPriorityQueue(arr3)
	pq3.top()
}

/*
ArrayList maxHeap:

コード内で、PriorityQueue 構造体は maxHeap という名前の整数スライスを持っています。このスライスは、優先度付きキューの要素を格納します。
NewPriorityQueue 関数では、引数として与えられた配列 arr から maxHeap スライスのディープコピーを作成しています。これにより、元の配列 arr は変更されずにそのコピーが変更されます。
そして、HeapLibraryのBuildMaxHeap 関数を使用して、この maxHeap 配列を最大ヒープに変換しています。
int top():

PriorityQueue 構造体には、top というメソッドが定義されています。このメソッドは、最大ヒープの最大値（ルート要素）を返します。
コードでは、top メソッドが正しく最大値をコンソールに出力し、その値を返していることが確認できます。
ただし、いくつかの改善点があります。

top メソッド内での fmt.Println の使用は、通常、ライブラリコード内では避けられます。なぜなら、このメソッドは値を返すべきであり、出力の責任はこのメソッドを呼び出す側にあります。
コードのエラーハンドリングが不足しています。たとえば、top メソッドは、ヒープが空の場合にインデックスエラーを引き起こす可能性があります。安全のため、このような場合に適切に対処する必要があります。
これらの点を考慮に入れると、コードは基本的な機能要件は満たしていますが、実際のプロダクション環境での使用を考慮すると、いくつかの改善が必要です。

*/
