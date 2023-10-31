//https://recursionist.io/dashboard/problems/450
/*
日本円を他の国の通貨に変換するツールを作成しています。convertJPYtoForeignCurrency(unit) は、設定した unit をクロージャ内に保存し、換算後の金額を出力する関数を返します。他の通貨への換算は、

"USD": 135
"EUR": 150
"GBP": 170
を使用してください。また、小数点以下切り捨てで計算してください。
*/
package main

import (
	"fmt"
)

func convertJPYtoForeignCurrency(unit string) func(int) {
	return func(yen int) {
		hashmap := map[string]int{
			"USD": 135,
			"EUR": 150,
			"GBP": 170,
		}
		foreign := int(yen / hashmap[unit])
		fmt.Printf("JPY: %d => %s: %d\n", yen, unit, foreign)
	}
}

func main() {
	doller := convertJPYtoForeignCurrency("USD")
	doller(10000)
	euro := convertJPYtoForeignCurrency("EUR")
	euro(10000)
	pound := convertJPYtoForeignCurrency("GBP")
	pound(10000)
}
