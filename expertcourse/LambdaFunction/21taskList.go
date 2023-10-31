// https://recursionist.io/dashboard/problems/449
// ToDo リストを文字列として格納した配列 toDoArr が与えられるので、ラムダクロージャを利用し、toDoArr の各要素を順番に出力してみましょう。
package main

import (
	"fmt"
)

func todoCall(arr []string) func() string {
	var count int = 0
	return func() string {
		if count == len(arr) {
			return "All done!"
		}
		result := arr[count]
		count++
		return result
	}
}

func main() {
	var todoArr []string = []string{"Read a Book", "Work out", "Recursion"}
	todoCaller := todoCall(todoArr)
	fmt.Println(todoCaller())
	fmt.Println(todoCaller())
	fmt.Println(todoCaller())
	fmt.Println(todoCaller())
}
