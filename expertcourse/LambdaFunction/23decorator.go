//https://recursionist.io/dashboard/problems/456
/*
デコレータ
medium
配列の合計値を返す sumOfArray 関数を受け取り、配列の各要素が 10 未満であればその個数とエラーメッセージを返し、全て条件を満たしていれば sumOfArray の結果を返す validateDecorator 関数を定義し、テストケースを出力するプログラムを作成してください。なお、validateDecorator 関数は、デコレータとして利用されるためのラッパー関数で、既存の sumOfArray 関数を受け取り、新しい機能である「エラーメッセージの出力」を付加してください。


*/
package main

import (
	"fmt"
)

func sumOfArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		if val >= 10 {
			sum += val
		}
	}
	return sum
}

func main() {
	validateDecorator := func(f func([]int) int) func([]int) {
		return func(arr []int) {
			errCount := 0
			for _, val := range arr {
				if val < 10 {
					errCount++
				}
			}
			if errCount == 0 {
				fmt.Printf("Sum of array is %d\n", f(arr))
			} else {
				fmt.Printf("%d error found\n", errCount)
			}
		}
	}

	sum := validateDecorator(sumOfArray)
	sum([]int{10, 20, 30, 40})
	sum([]int{9, 10, 20, 30})
	sum([]int{3, 5, 40, 50})
}
