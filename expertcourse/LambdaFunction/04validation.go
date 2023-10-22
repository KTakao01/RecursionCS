//https://recursionist.io/dashboard/problems/428
/*Clara は正しい Email を入力されているかチェックし、ユーザに Error メッセージを返す関数を作成しています。それでは、文字列を検証する validator を 3 つ作成し、指定の Message を返す emailValidation(bool(string), string) を作成して下さい。validator は 下記 3 つです。

bool doesNotStartWithAt(string email)
文字列の先頭が ＠ から始まらない場合 true を返し、始まる場合 false を返すラムダ関数

bool doesNotHaveSpace(string email)
文字列に空白が含まれない場合 true を返し、含まれる場合 false を返すラムダ関数

bool hasUppercaseAndLowercase(string email)
文字列に大文字と小文字を含む場合 true を返し、含まない場合 false を返すラムダ関数
*/

package main

import (
	"fmt"
	"reflect"
	"strings"
)

type stringToBoolFunc func(string) bool

func main() {
	emailValidation := func(f stringToBoolFunc, email string) string {
		if f(email) {
			return "Email is correct."
		}
		return "Email is not correct."
	}

	doesNotStartWithAt := func(email string) bool {
		if email[0:1] != "@" {
			return true
		}
		return false
	}

	doesNotHaveSpace := func(email string) bool {
		for i := 0; i < len(email); i++ {
			if email[i:i+1] == " " {
				return false
			}
		}
		return true
	}
	hasUppercaseAndLowercase := func(email string) bool {
		if !(reflect.DeepEqual(strings.ToUpper(email), email) || reflect.DeepEqual(strings.ToLower(email), email)) {
			return true
		}
		return false

	}

	fmt.Println(emailValidation(doesNotStartWithAt, "@gmail.com"))
	fmt.Println(emailValidation(doesNotStartWithAt, "kkk@gmail.com"))
	fmt.Println(emailValidation(doesNotHaveSpace, "Hello world"))
	fmt.Println(emailValidation(doesNotHaveSpace, "Helloworld"))
	fmt.Println(emailValidation(hasUppercaseAndLowercase, "hello world"))
	fmt.Println(emailValidation(hasUppercaseAndLowercase, "HELLO WORLD"))
	fmt.Println(emailValidation(hasUppercaseAndLowercase, "Hello world"))

}
