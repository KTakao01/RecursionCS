package main

import (
    "fmt"
    "strings"
)
type Item struct{
    data int32
    next *Item
}

func NewItem(data int32)*Item{
    return &Item{data:data}
}

type Stack struct{
    head *Item
}

func NewStack() *Stack{
    return &Stack{}
}

func (s *Stack) push(data int32){
    temp := s.head
    newHead := NewItem(data)
    s.head = newHead
    s.head.next = temp
}

func (s *Stack) pop() *int32{
    temp := s.head
    s.head = s.head.next
    fmt.Println(temp.data)
    return &temp.data
}

func (s *Stack) peek() *int32{
    fmt.Println(s.head.data)
    return &s.head.data
}

func reverse(arr []int32) []int32{
    // 関数を完成させてください
    s := NewStack()
    for i := 0; i < len(arr); i++{
        s.push(arr[i])
    }
    reverseArr := make([]int32,len(arr))
    for i := 0; i < len(arr); i++{
        reverseArr[i] = *s.pop()
    }
    return reverseArr
}