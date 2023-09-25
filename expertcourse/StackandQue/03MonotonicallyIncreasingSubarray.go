package main

import (
	"fmt"
)

type Item struct {
	data int32
	next *Item
}

func NewItem(data int32) *Item {
	return &Item{data: data}
}

type Stack struct {
	head *Item
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(data int32) {
	temp := s.head
	newHead := NewItem(data)
	s.head = newHead
	s.head.next = temp
}

func (s *Stack) pop() *int32 {
	if s.head == nil {
		return nil
	}
	temp := s.head
	s.head = s.head.next
	fmt.Println(temp.data)
	return &temp.data
}

func (s *Stack) peek() *int32 {
	if s.head == nil {
		return nil
	}
	fmt.Println(s.head.data)
	return &s.head.data
}

func consecutiveWalk(arr []int32) []int32 {
	// 関数を完成させてください
	if len(arr) == 0 {
		return []int32{}
	}
	s := NewStack()
	s.push(arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] <= *s.peek() {
			for s.peek() != nil {
				s.pop()
			}
		}
		s.push(arr[i])
	}
	var result []int32
	for s.peek() != nil {
		result = append(result, []int32{*s.pop()}...)
	}
	return result
}

//空のスタックを用意
//配列の先頭とスタックの最初の要素を比較して、スタックの最初<配列の先頭の時に配列をスタックにpushする。スタックが空でもpushする。
//配列>=スタックとなったときはスタックをpopしてリセット後に新しくその配列をpushしたスタックを作成する。上記を繰り返す。
//最後に残ったスタックを部分単調増加となっている
