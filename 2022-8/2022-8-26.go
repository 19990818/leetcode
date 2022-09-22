package main

type MyQueue struct {
	stack1 stack
	stack2 stack
}

func Constructor2() MyQueue {
	return MyQueue{
		stack1: stack{make([]int, 0)},
		stack2: stack{make([]int, 0)},
	}
}

func (this *MyQueue) Push(x int) {
	this.stack1.push(x)
}

func (this *MyQueue) Pop() int {
	if this.stack2.empty() {
		for !this.stack1.empty() {
			this.stack2.push(this.stack1.pop())
		}
	}
	return this.stack2.pop()
}

func (this *MyQueue) Peek() int {
	if this.stack2.empty() {
		for !this.stack1.empty() {
			this.stack2.push(this.stack1.pop())
		}
	}
	return this.stack2.peek()
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1.arr) == 0 && len(this.stack2.arr) == 0
}

type stack struct {
	arr []int
}

func (this *stack) pop() int {
	res := this.arr[len(this.arr)-1]
	this.arr = this.arr[0 : len(this.arr)-1]
	return res
}
func (this *stack) push(val int) {
	this.arr = append(this.arr, val)
}
func (this *stack) empty() bool {
	return len(this.arr) == 0
}
func (this *stack) peek() int {
	return this.arr[len(this.arr)-1]
}
