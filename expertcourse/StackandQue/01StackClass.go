package main

import "fmt"

type Item struct{
    data int
    next *Item
}

func NewItem(data int) *Item{
    return &Item{data:data}
}

type Stack struct{
    head *Item
}

func NewStack() *Stack{
    return &Stack{head:nil}
}

func(s *Stack) Push(data int){
    newNode := NewItem(data)
    temp := s.head
    s.head = newNode
    s.head.next = temp
}

func(s *Stack) Pop()*int{
    if s.head == nil{
        return nil
    }
    temp := s.head
    s.head = s.head.next
    fmt.Println(temp.data)
    return &temp.data
}

func(s *Stack) Peek()*int{
    fmt.Println(s.head.data)
    return &s.head.data
}


func main() {
    s1 := NewStack()
    s1.Push(2)
    s1.Peek()
    s1.Push(4)
    s1.Push(3)
    s1.Push(1)
    s1.Pop()
    s1.Peek()

    s2 := NewStack()
    s2.Pop()
    s2.Push(9)
    s2.Push(8)
    s2.Peek()
}
