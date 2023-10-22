/*
https://recursionist.io/dashboard/problems/419
自然数 n が与えられるので、n 個の 0 を文字列として追加し、その文字列を返す、duplicateZero のラムダを作成し、テストケースを出力してください。
*/
package main

import (
	"fmt"
)

func main() {
	duplicateZero := func(x int) string {
		results := ""
		for i := 1; i <= x; i++ {
			results += "0"
		}
		return results
	}

	fmt.Println(duplicateZero(5))

	fmt.Println(duplicateZero(10))
}
