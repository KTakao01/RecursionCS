// https://recursionist.io/dashboard/problems/437
package main

import (
	"fmt"
)

func majority(f func(int) bool, arr []int) bool {
	count := 0
	for i := 0; i < len(arr); i++ {
		if f(arr[i]) {
			count++
		}
	}
	middle := len(arr) / 2
	if count > middle {
		return true
	}
	return false
}

func isOdd(n int) bool {
	if n%2 != 0 {
		return true
	}
	return false
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(majority(isOdd, []int{1, 2, 3, 4, 5}))
	fmt.Println(majority(isOdd, []int{2, 4, 6, 7, 8}))
	fmt.Println(majority(isEven, []int{3, 6, 8, 12, 15}))
	fmt.Println(majority(isEven, []int{4, 5, 7, 11, 14}))
}
