package golang

type MyQueue struct {
	queueList []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.queueList = append(this.queueList, x)
}

func (this *MyQueue) Pop() int {
	firstInt := 0
	firstInt = this.queueList[0]
	this.queueList = this.queueList[1:]
	return firstInt
}

func (this *MyQueue) Peek() int {
	return this.queueList[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.queueList) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
