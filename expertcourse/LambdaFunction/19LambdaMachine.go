//https://recursionist.io/dashboard/problems/447
/*
加算や乗算を行うラムダ関数と文字列を入力として取り込み、管理し、出力として返す Lambda Machine というデータ構造体を作成してください。例えば、"addition"という文字列と、加算を表すラムダ関数 addition を入力し、それらを handler に格納することができます。使用するときは、"addition" という文字列を用いて呼び出すことで、加算を実行することができます。

lambda Machine は以下のデータ構造を持ちます。

HashMap handlers
ハンドラはキーと値のペアを含み、キーは関数名、値は ラムダ関数とします。これによって handlers[key] で特定のラムダを取得することができます。

void insert(String key, int(int, int))
関数名を表すキーと、関数を実行するラムダ関数をペアとしてデータ構造に挿入します。

LambdaFunction retrieve(string key)
キー（関数名）を受け取り、ハッシュマップにペアとしてマップされたラムダ関数を返します。

以下の整数 x、y を受け取り特定の値を返すラムダ関数を作成して、データ構造に挿入してみましょう。

LambdaFunction pythagora(int x, int y)
直角三角形の 2 辺の長さを x、y とし、斜辺の長さを返すラムダ。出力は小数点を切り捨ててください。

LambdaFunction addition(int x, int y)
x と y を足した値を返すラムダ。

LambdaFunction multiplication(int x, int y)
x と y を掛けた値を返すラムダ。
*/
package main

import (
	"fmt"
	"math"
)

type lambdaMachine struct {
	handlers map[string]func(int, int) int
}

func (lm *lambdaMachine) insert(key string, f func(int, int) int) {
	if lm.handlers == nil {
		lm.handlers = make(map[string]func(int, int) int)
	}
	if _, ok := lm.handlers[key]; !ok {
		lm.handlers[key] = f
	} else {
	}
}

func (lm *lambdaMachine) retrieve(key string) func(int, int) int {
	return lm.handlers[key]
}

func main() {
	pythagora := func(x, y int) int {
		return int(math.Pow(float64(x*x+y*y), 0.5))
	}

	lm := lambdaMachine{}
	lm.insert("pythagora", pythagora)
	fmt.Println(lm.retrieve("pythagora")(3, 4))

	addition := func(x, y int) int {
		return x + y
	}

	lm.insert("addition", addition)
	fmt.Println(lm.retrieve("addition")(2, 5))

	multiplication := func(x, y int) int {
		return x * y
	}

	lm.insert("multiplication", multiplication)
	fmt.Println(lm.retrieve("multiplication")(4, 10))

}
