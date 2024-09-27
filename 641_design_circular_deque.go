package main

type MyCircularDeque struct {
	buff  []int
	size  int
	front int
	rear  int
}

func Constructor1(k int) MyCircularDeque {
	return MyCircularDeque{
		buff:  make([]int, k),
		size:  k,
		front: -1,
		rear:  -1,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.front = 0
		this.rear = 0
	} else {
		this.front = (this.front - 1 + this.size) % this.size
	}
	this.buff[this.front] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.front = 0
		this.rear = 0
	} else {
		this.rear = (this.rear + 1) % this.size
	}
	this.buff[this.rear] = value
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if this.front == this.rear { // the only element
		this.front = -1
		this.rear = -1
	} else {
		this.front = (this.front + 1) % this.size
	}
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.front == this.rear {
		this.front = -1
		this.rear = -1
	} else {
		this.rear = (this.rear - 1 + this.size) % this.size
	}
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buff[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buff[this.rear]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.front == -1
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.rear+1)%this.size == this.front
}
