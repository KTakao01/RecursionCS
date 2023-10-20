/*
https://recursionist.io/dashboard/problems/452
問題 451 では、二分ヒープを構築し、最も優先度の高い値にアクセスできるようになりました。それでは、最も優先度の高い要素を削除する pop() メソッドを実装してみましょう。

ArrayList maxHeap
動的配列で、優先度付きキューの要素を保持します。この配列は最大ヒープと呼ばれる特殊な形状に整列します。これにより、最大の要素が常に根に存在することが保証されます。arr が与えられたとき、その状態を変更せずにこの配列を最大ヒープに変換する必要があります。arr のディープコピーを作成し元の配列の変更を防ぎつつ、HeapLibrary の buildMaxHeap 関数を使用してそのコピーを最大ヒープに変換することで達成できます。

int top()
最大ヒープの最大値を返します。

integer pop()
優先度付きキューから最も優先度の高い値を削除して返します。この操作を実行する際には、ヒープのプロパティ（最大ヒープの場合は親ノードの値が子ノードの値よりも大きい）が維持されるように注意する必要があります。最初に、根ノードを取り出し、ヒープの最後の葉ノードと入れ替え、この新しい根ノードに対して maxHeapify を呼び出すことで、実現できます。
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

func (pq *PriorityQueue) pop() int {
	pq.maxHeap[0], pq.maxHeap[len(pq.maxHeap)-1] = pq.maxHeap[len(pq.maxHeap)-1], pq.maxHeap[0]
	top := pq.maxHeap[len(pq.maxHeap)-1]
	pq.maxHeap = pq.maxHeap[:len(pq.maxHeap)-1]
	heapLib := HeapLibrary{}
	heapLib.MaxHeapify(pq.maxHeap, len(pq.maxHeap)-1, 0)
	fmt.Println(top)
	return top
}

func main() {
	arr1 := []int{2, 3, 43, 2, 53, 6, 75, 10}
	pq1 := NewPriorityQueue(arr1)
	pq1.pop()
	pq1.pop()
	pq1.pop()
	pq1.pop()
}
