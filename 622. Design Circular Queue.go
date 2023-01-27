package golang

type MyCircularQueue struct {
	myQueue   []int
	size      int
	count     int
	headIndex int
}

func Constructor(k int) MyCircularQueue {
	myQueue := make([]int, k)
	return MyCircularQueue{myQueue, k, 0, 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.myQueue[(this.headIndex+this.count)%this.size] = value
	this.count++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.headIndex = (this.headIndex + 1) % this.size
	this.count--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.myQueue[this.headIndex]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.myQueue[(this.headIndex+this.count-1)%this.size]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.count == this.size
}
