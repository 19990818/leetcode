package main

import "fmt"

type AuthenticationManager struct {
	arr    []int
	m      map[string]int
	tolive int
}

func ConstructorAuth(timeToLive int) AuthenticationManager {
	return AuthenticationManager{make([]int, 0), make(map[string]int), timeToLive}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.m[tokenId] = currentTime
	if len(this.arr) >= this.tolive {
		this.arr = this.arr[1:]
	}
	this.arr = append(this.arr, currentTime)
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if _, ok := this.m[tokenId]; ok {
		if this.m[tokenId]+this.tolive > currentTime {
			this.m[tokenId] = currentTime
		} else {
			delete(this.m, tokenId)
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	ans := 0
	for _, val := range this.m {
		if currentTime-val < this.tolive {
			ans++
		}
	}
	return ans
}

type MyLinkedList struct {
	list *ListNode
	size int
}

func ConstructorList() MyLinkedList {
	return MyLinkedList{nil, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.size || index < 0 {
		return -1
	}
	// printList(this.list)
	cur := this.list
	for index > 0 {
		cur = cur.Next
		//fmt.Println(cur.Val)
		index--
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	res := new(ListNode)
	res.Val = val
	res.Next = this.list
	this.list = res
	this.size++
	//printList(this.list)
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.list
	if cur == nil {
		node := new(ListNode)
		node.Val = val
		this.list = node
		this.size++
		return
	}
	for cur.Next != nil {
		cur = cur.Next
	}
	node := new(ListNode)
	node.Val = val
	cur.Next = node
	this.size++
	// printList(this.list)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	cur := this.list
	for index > 1 {
		cur = cur.Next
		index--
	}
	tempPost := cur.Next
	newNode := new(ListNode)
	newNode.Val = val
	newNode.Next = tempPost
	cur.Next = newNode
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.size <= index || index < 0 {
		return
	}
	if index == 0 {
		this.list = this.list.Next
		this.size--
		return
	}
	cur := this.list
	for index > 1 {
		cur = cur.Next
		index--
	}
	cur.Next = cur.Next.Next
	this.size--
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val, " ")
		list = list.Next
	}
	fmt.Println()
}

type MyCircularQueue struct {
	queue []int
	head  int
	tail  int
	count int
}

func ConstructorCir(k int) MyCircularQueue {
	temp := make([]int, k)
	for i := 0; i < k; i++ {
		temp[i] = -1
	}
	return MyCircularQueue{temp, 0, 0, 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.count >= len(this.queue) {
		return false
	}
	this.count++
	this.queue[this.tail] = value
	this.tail = (this.tail + 1) % len(this.queue)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.count <= 0 {
		return false
	}
	this.count--
	this.head = (this.head + 1) % len(this.queue)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.count == 0 {
		return -1
	}
	return this.queue[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.count == 0 {
		return -1
	}
	return this.queue[(this.tail+len(this.queue)-1)%len(this.queue)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.count == len(this.queue)
}
