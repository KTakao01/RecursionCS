/*
https://recursionist.io/dashboard/problems/430
自然数 n 以下に含まれる、以下を計算する関数を作成しています。

奇数の総和
3 と 5 の倍数の総和
素数の総和
自然数 n と、以下のラムダ関数を受け取り、n 以下の特定の値の合計値を返す summation(bool(int), int) を作成してください。

奇数かどうか判定する関数
bool isOdd(int x)

3 または 5 の倍数かどうか判定する関数
bool isMultipleOf3Or5(int x)

素数かどうか判定する関数
bool isPrime(int x)
*/

package main

import (
	"fmt"
)

type intToBoolFunc func(int) bool

func main() {
	summation := func(f intToBoolFunc, n int) int {
		var sum int
		for x := 1; x <= n; x++ {
			if f(x) {
				sum += x
			}
		}
		return sum
	}

	isOdd := func(x int) bool {
		if x%2 == 1 {
			return true
		}
		return false
	}

	isMultipleOf3Or5 := func(x int) bool {
		if x%3 == 0 || x%5 == 0 {
			return true
		}
		return false
	}

	isPrime := func(x int) bool {
		if x <= 1 {
			return false
		}
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}

	fmt.Println(summation(isOdd, 3))
	fmt.Println(summation(isOdd, 10))
	fmt.Println(summation(isOdd, 25))
	fmt.Println(summation(isMultipleOf3Or5, 3))
	fmt.Println(summation(isMultipleOf3Or5, 10))
	fmt.Println(summation(isMultipleOf3Or5, 100))
	fmt.Println(summation(isPrime, 2))
	fmt.Println(summation(isPrime, 10))
	fmt.Println(summation(isPrime, 100))
}
