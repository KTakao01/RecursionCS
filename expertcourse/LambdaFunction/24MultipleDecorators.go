//https://recursionist.io/dashboard/problems/457
/*
任意の関数に機能を追加するデコレータをいくつか作成し、テストケースを出力してみましょう。['h','E','L','L','O'] のような文字によって構成された配列 arr を、"hELLO" のような文字列に変換して返す toString という関数、2 つの文字列を連結する concat という関数を作成し、以下のデコレータに渡すことで機能を追加してください。

LambdaFunction capitalizeDecorator(string(char[]))
「文字の配列を受け取る関数の参照を受け取り、入力の配列の先頭を大文字にして文字列を返す機能」を追加するラムダ

LambdaFunction lowercaseResultDecorator(string(char[]))
「文字の配列受け取る関数の参照を受け取り、小文字へ変換して文字列を返す機能」を追加するラムダ。

LambdaFunction printInfoDecorator(string(string, string))
「文字の配列受け取る関数の参照を受け取り、引数を連結する機能」を追加するラムダ。

ラムダに渡す関数は以下を参考にしてください。

string toString(char[] arr)
文字の配列を受け取り、文字列として連結して返す関数。

string concat(string s1, string s2)
文字列を受け取り、連結して返す関数。
*/
package main

import (
	"fmt"
	"strings"
)

func toString(arr []rune) string {
	var result string
	for i := 0; i < len(arr); i++ {
		result += string(arr[i])
	}
	return result
}

func concat(str1, str2 string) string {
	return str1 + str2
}

func main() {
	capitalizeDecorator := func(f func([]rune) string) func([]rune) string {
		return func(arr []rune) string {
			str := f(arr)

			return strings.ToUpper(str[0:1]) + str[1:]
		}
	}

	arr := []rune{'h', 'E', 'L', 'L', 'O'}
	f1 := capitalizeDecorator(toString)
	fmt.Println(f1(arr))

	lowercaseResultDecorator := func(f func([]rune) string) func([]rune) string {
		return func(arr []rune) string {
			str := f(arr)
			return strings.ToLower(str)
		}
	}

	f2 := lowercaseResultDecorator(toString)
	fmt.Println(f2(arr))

	printInfoDecorator := func(f func(string, string) string) func(string, string) string {
		return func(str1, str2 string) string {
			str := f(str1, str2)
			return fmt.Sprintf("%s + %s = %s\n", str1, str2, str)
		}
	}

	f3 := printInfoDecorator(concat)
	fmt.Println(f3(f1(arr), f2(arr)))
}
