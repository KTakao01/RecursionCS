/*
問題 447 では、ラムダを自由に出し入れすることができるデータ構造 Lambda Machine を作成しました。今回は、ラウンドロビンによってラムダを返す roundRobinRetrieve() メソッドを定義します。これは、ラムダ関数を格納した配列を管理する lambdaStorage という配列を追加し、その配列からラムダ関数を順番に取り出すことで実現します。

また、配列から特定の文字列を指定できるように handler には、ラムダ関数を識別するための文字列と配列のインデックスをペアとして格納します。

なお、ラウンドロビンとは、特定のデータ項目を順番に処理する方式のことです。通常は、項目を順番に取り出して処理し、最後まで処理したら最初から再度処理するというように繰り返し行われます。

LambdaFunction roundRobinRetrieve()
ルックアップテーブルに基づいて関数を返す場合とは別に、ラウンドロビンによって返す関数を選択します。

https://recursionist.io/dashboard/problems/448

*/

package main

import (
	"fmt"
	"math"
)

type lambdaMachine struct {
	handlers      map[string]int
	lambdaStorage []func(int, int) int
	currentIndex  int
}

//handlers[] = index lambdaStorage[index] = func

func (lm *lambdaMachine) insert(key string, f func(int, int) int) {
	if lm.handlers == nil {
		lm.handlers = make(map[string]int)
		lm.lambdaStorage = []func(int, int) int{}
	}

	if _, ok := lm.handlers[key]; !ok {
		index := len(lm.lambdaStorage)
		lm.lambdaStorage = append(lm.lambdaStorage, f)
		lm.handlers[key] = index
	} else {
	}
}

func (lm *lambdaMachine) retrieve(key string) func(int, int) int {
	return lm.lambdaStorage[lm.handlers[key]]
}

// バッグ：len(lm.lamdaStorage),*1ずつfunc入っている,
func (lm *lambdaMachine) roundRobinRetrieve() func(int, int) int {
	maxIndex := len(lm.lambdaStorage)
	result := lm.lambdaStorage[lm.currentIndex]
	lm.currentIndex = (lm.currentIndex + 1) % maxIndex
	return result
}

func main() {
	pythagora := func(x, y int) int {
		return int(math.Pow(float64(x*x+y*y), 0.5))
	}

	lm := lambdaMachine{}
	lm.insert("pythagora", pythagora)

	addition := func(x, y int) int {
		return x + y
	}

	lm.insert("addition", addition)

	multiplication := func(x, y int) int {
		return x * y
	}

	lm.insert("multiplication", multiplication)
	func1 := lm.roundRobinRetrieve()(6, 8)
	func2 := lm.roundRobinRetrieve()(6, 8)
	func3 := lm.roundRobinRetrieve()(6, 8)
	func4 := lm.roundRobinRetrieve()(6, 8)
	fmt.Println(func1)
	fmt.Println(func2)
	fmt.Println(func3)
	fmt.Println(func4)
}
