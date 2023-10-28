/*
https://recursionist.io/dashboard/problems/417
円の半径 radius が与えられるので、円の面積を返す、circleArea のラムダを作成し、テストケースを出力してください。ただし円周率は 3.14 とします。
*/
package main

import (
	"fmt"
)

func main() {
	circleArea := func(x int) float64 { return float64(x*x) * 3.14 }
	fmt.Println(circleArea(1))
	fmt.Println(circleArea(5))
}
