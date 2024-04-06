package main

type MyCircularQueue struct {
	data  []int
	first int
	last  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data:  make([]int, k),
		first: -1,
		last:  -1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.first = 0
	}

	this.last = (this.last + 1) % len(this.data)
	this.data[this.last] = value

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.first == this.last {
		this.first = -1
		this.last = -1
	} else {
		this.first = (this.first + 1) % len(this.data)
	}
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.first]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.last]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.first == -1 //&& this.last == -1
}

func (this *MyCircularQueue) IsFull() bool {
	if this.last >= this.first {
		return this.first == 0 && this.last == len(this.data)-1
	}
	return this.last+1 == this.first
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
