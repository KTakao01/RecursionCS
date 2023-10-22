/*
https://recursionist.io/dashboard/problems/429
動物の年齢による医療保険の加入可否を決める関数を作成しています。人間の年齢に換算して 60 歳以下であれば true, 60 歳より上であれば false を返す qualifiedForInsurance(int(int), int) を作成してください。動物の年齢から人の年齢に換算するラムダ関数は、下記 2 つです。

int dogToHuman(int dogAge)
犬の年齢を人の年齢に変換する関数。人間の年齢への換算は、20 + (犬年齢 - 2）* 7 で求めてください。

int catToHuman(int catAge)
猫の年齢を人の年齢に変換する関数。人間の年齢への換算は、24 + (猫年齢 - 2) * 4 で求めてください。
*/
package main

import (
	"fmt"
)

type intTointFunc func(int) int

func main() {
	qualifiedForInsurance := func(f intTointFunc, x int) bool {
		if f(x) > 60 {
			return false
		}
		return true
	}

	dogToHuman := func(dogAge int) int {
		return 20 + (dogAge-2)*7
	}
	catToHuman := func(catAge int) int {
		return 24 + (catAge-2)*4
	}

	fmt.Println(qualifiedForInsurance(dogToHuman, 7))
	fmt.Println(qualifiedForInsurance(dogToHuman, 8))
	fmt.Println(qualifiedForInsurance(catToHuman, 11))
	fmt.Println(qualifiedForInsurance(catToHuman, 12))

}
