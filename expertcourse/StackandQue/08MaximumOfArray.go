package main

type Item struct {
	data int32
	prev *Item
	next *Item
}

type Dequeue struct {
	head *Item
	tail *Item
}

func NewDequeue() *Dequeue {
	return &Dequeue{}
}

func (d *Dequeue) peekFront() *int32 {
	if d.head == nil {
		return nil
	}
	//fmt.Println(d.head.data)
	return &d.head.data
}

func (d *Dequeue) enqueueFront(data int32) {
	newNode := &Item{data: data}
	if d.head == nil {
		d.head = newNode
		d.tail = d.head
	}
	newNode.prev = d.head
	d.head.next = newNode
	d.head = newNode
}

func (d *Dequeue) enqueueBack(data int32) {
	newNode := &Item{data: data}
	if d.tail == nil {
		d.tail = newNode
		d.head = d.tail
	}
	newNode.prev = d.tail
	d.tail.next = newNode
	d.tail = newNode
}

func getMax(arr []int32) int32 {
	// 関数を完成させてください
	d := NewDequeue()
	d.enqueueFront(arr[0])
	for i := 0; i < len(arr); i++ {
		if *d.peekFront() < arr[i] {
			d.enqueueFront(arr[i])
		} else {
			d.enqueueBack(arr[i])
		}
	}
	return *d.peekFront()

}
