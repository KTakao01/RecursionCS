// https://recursionist.io/dashboard/problems/435
package main

import (
	"fmt"
)

func maxByCriteria(f func(string, string) bool, str []string) {
	m1 := str[0]
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str); j++ {
			//fmt.Println("比較前i,j,max",i,j,m1)
			if i != j {
				if f(str[i], str[j]) && f(str[i], m1) {
					m1 = str[i]
					//fmt.Println("比較結果i,j,max",i,j,m1)
				}
			}
		}
	}
	fmt.Println(m1)

}

func compareLength(s1, s2 string) bool {
	if len(s1) >= len(s2) {
		return true
	}
	return false
}

func compareAsciiTotal(s1, s2 string) bool {
	aS1 := []rune(s1)
	aS2 := []rune(s2)

	var r1, r2 int
	for _, val := range aS1 {
		r1 += int(val)
	}

	for _, val := range aS2 {
		r2 += int(val)
	}
	if r1 >= r2 {
		return true
	}

	return false
}

func main() {
	maxByCriteria(compareLength, []string{"apple", "yumberry", "grape", "banana", "mandarin"})
	maxByCriteria(compareLength, []string{"zoomzoom", "choochoo", "beepbeep", "ahhhahhh"})
	maxByCriteria(compareAsciiTotal, []string{"apple", "yumberry", "grape", "banana", "mandarin"})
	maxByCriteria(compareAsciiTotal, []string{"zoom", "choochoo", "beepbeep", "ahhhahhh"})

}
