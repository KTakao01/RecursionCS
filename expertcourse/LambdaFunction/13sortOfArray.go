// https://recursionist.io/dashboard/problems/438
package main

import (
	"fmt"
	"sort"
)

func sortByAscii(arr []string) []string {
	sort.SliceStable(arr, func(i, j int) bool {
		return compareByAscii(arr[i], arr[j])
	})
	return arr
}

func compareByAscii(a, b string) bool {
	//a < b: true //a >= b: false
	iA := changeToRune(a)
	iB := changeToRune(b)
	fmt.Println("比較", a, iA, b, iB)
	return iA < iB
}

func changeToRune(str string) int {
	result := 0
	for _, val := range str {
		result += int(val)
	}
	return result
}
