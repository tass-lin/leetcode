package main

type MyStack struct {
	val []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{[]int{}}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.val = append([]int{x},this.val...)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	temp := this.val[0]
	this.val = this.val[1:]
	return temp
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.val[0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.val) > 0 {
		return false
	}
	return true
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

//MyStack stack = new MyStack();
//
//stack.push(1);
//stack.push(2);
//stack.top();   // returns 2
//stack.pop();   // returns 2
//stack.empty(); // returns false