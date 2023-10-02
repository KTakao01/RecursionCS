package main

import (
	"fmt"
)

func minWindowArrK(intArr []int32, k int32) []int32 {
	// 関数を完成させてください
	results := []int32{}
	deque := []int32{}
	//最初の要素kをキューに追加する
	for i := int32(0); i < k; i++ {
		current := intArr[i]
		//スタックが空でないときに配列要素が末尾要素より小さいなら、大きくなるまで末尾要素をdeque
		for len(deque) != 0 && intArr[deque[len(deque)-1]] >= current {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	fmt.Println(deque)

	//ウィンドウサイズを一つずらした後の処理
	for i := k; i < int32(len(intArr)); i++ {
		results = append(results, intArr[deque[0]])
		current := intArr[i]
		for len(deque) != 0 && intArr[deque[len(deque)-1]] >= current {
			deque = deque[:len(deque)-1]
			fmt.Println("a", deque)
		}
		//ウィンドウサイズの範囲外の要素を削除
		if len(deque) != 0 && deque[0] == i-k {
			deque = deque[1:len(deque)]
			fmt.Println("b", deque)
		}

		deque = append(deque, i)
		fmt.Println("c", deque)
	}
	results = append(results, intArr[deque[0]])
	fmt.Println("d", deque)
	return results
}
