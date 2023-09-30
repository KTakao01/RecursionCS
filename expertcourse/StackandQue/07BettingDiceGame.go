package main

import (
	"fmt"
	"strconv"
)

func diceStreakGamble(player1 []int32, player2 []int32, player3 []int32, player4 []int32) string {
	// 関数を完成させてください
	r1 := consectiveWalk(player1)
	r2 := consectiveWalk(player2)
	r3 := consectiveWalk(player3)
	r4 := consectiveWalk(player4)

	p1 := wonPrize(r1)
	p2 := wonPrize(r2)
	p3 := wonPrize(r3)
	p4 := wonPrize(r4)

	var hashmap map[int32][]int32
	var max int32 = p1
	hashmap = map[int32][]int32{p1: r1}
	var mIndex int32 = 1
	if max < p2 {
		max = p2
		mIndex = 2
		hashmap = map[int32][]int32{p2: r2}
	}
	if max < p3 {
		max = p3
		mIndex = 3
		hashmap = map[int32][]int32{p3: r3}
	}
	if max < p4 {
		max = p4
		mIndex = 4
		hashmap = map[int32][]int32{p4: r4}
	}
	fmt.Println(max, mIndex, hashmap)
	return fmt.Sprintf("Winner: Player %d won $%d by rolling %s", mIndex, max, printArray(hashmap[max]))
}

func consectiveWalk(player []int32) []int32 {
	var stack []int32
	//初期の配列要素をpushする
	stack = append(stack, player[0])

	//スタックを降順にする
	//配列要素とスタック末尾を比較して配列要素が大きければ空になるかスタック末尾が大きくなるまでスタックをpopして末尾に配列要素を挿入。
	//配列要素が等しいか小さければそのまま末尾に挿入。
	for i := 1; i < len(player); i++ {
		if stack[len(stack)-1] > player[i] {
			for len(stack) != 0 {
				fmt.Println(stack, player[i])
				stack = stack[:len(stack)-1]
				fmt.Println(stack)
			}
		}
		stack = append(stack, player[i])
	}
	fmt.Println("最終", stack)
	return stack
}

func wonPrize(rolling []int32) int32 {
	var ans int32
	ans = int32(len(rolling)) * 4
	return ans
}

func printArray(arr []int32) string {
	var result string = "["
	for i := 0; i < len(arr); i++ {
		str := strconv.Itoa(int(arr[i]))
		result += str

		if i == len(arr)-1 {
			result += "]"
		} else {
			result += ","
		}
	}
	return result
}

//キーワード：リセット、単調増加→スタック
//スタックに残った単調増加部分配列の要素数✖️$4がプレイヤーが得られる値。

//最大値は総当たりを繰り返す？
//最大値を更新していけば１回の総当たりで済む
