package main

type MyQueue struct {
	val []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{[]int{}}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.val = append(this.val,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	temp := this.val[0]
	this.val = this.val[1:]
	return temp
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.val[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.val) > 0 {
		return false
	}
	return true
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

//MyQueue queue = new MyQueue();
//
//queue.push(1);
//queue.push(2);
//queue.peek();  // returns 1
//queue.pop();   // returns 1
//queue.empty(); // returns false

func main() {
	
}
