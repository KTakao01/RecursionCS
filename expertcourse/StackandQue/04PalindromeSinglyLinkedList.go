package main

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type Stack struct {
	head *SinglyLinkedListNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(data int32) {
	temp := s.head
	s.head = &SinglyLinkedListNode{data: data}
	s.head.next = temp
}

func (s *Stack) Pop() *int32 {
	if s.head == nil {
		return nil
	}
	temp := s.head
	s.head = s.head.next
	return &temp.data
}

func (s *Stack) Peek() *int32 {
	if s.head == nil {
		return nil
	}
	return &s.head.data
}

func palindromeLinkedList(head *SinglyLinkedListNode) bool {
	// 関数を完成させてください
	fast := head
	slow := head
	s := NewStack()

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		s.Push(slow.data)
		slow = slow.next
	}
	fmt.Println(fast)
	fmt.Println(slow, *s.Peek())
	//ノード数が奇数のとき
	if fast != nil {
		slow = slow.next
	}

	for slow != nil {
		if *s.Peek() != slow.data {
			fmt.Println(slow.data, *s.Peek())
			return false
		}
		s.Pop()
		slow = slow.next
	}
	return true
}

//問題文
//連結リストの先頭 head が与えられるので、それが回文になっているか判定する、palindromeLinkedList という関数を作成してください。

//スタックを使わない方法その１
//len k.  len arr -1 - k//len 2 5-2 =3
//fast とslowつかえば　　先にfastだけk進む　このfastをfirstとして保存
//slowとfastを進ませてfastが末尾にきたら止める//arr-1 -kだけfastすすむのでslowはarr-1-kの位置にいる
//このslowをsecondとして保存
//fast と secondが一致していればいい。
//二重ループにしてk=0;arr-1-k > k;k++まで進める。
//２重ループの中にforループが２回入る。３重ループと計算量が多い。

//スタックを使わない方法その２
//fast,slowで真ん中のノードまで
//真ん中のノードにきたところで、slowをリストの終わりまでを反転する
//slowは反転させたノードの先頭をさすので、fastをリセットしてslowとfastを比較
//反転するのでひとつ余分にループがいる。

//スタックを使う方法
//反転をスタック構造で表すことを考える。
//先頭から真ん中までを反転させられる（片方向リストではノードを進める必要があるので反転したいところを起点として、つまり真ん中から最後まで）
//
