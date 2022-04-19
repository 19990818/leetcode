package main

import (
	"container/heap"
	"math"
	"sort"
)

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	ans := math.MaxInt32
	for i := 0; i < len(nums)-1; i++ {
		a := nums[i] + k
		b := nums[len(nums)-1] - k
		maxNum := max(a, b)
		c := nums[0] + k
		d := nums[i+1] - k
		minNum := min(c, d)
		ans = min(maxNum-minNum, ans)
	}
	ans = min(ans, nums[len(nums)-1]-nums[0])
	return ans
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1, cur2 := l1, l2
	res := new(ListNode)
	ans := res
	c := 0
	for cur1 != nil && cur2 != nil {
		ans.Val = (cur1.Val + cur2.Val + c) % 10
		c = (cur1.Val + cur2.Val + c) / 10
		cur1 = cur1.Next
		cur2 = cur2.Next
		if cur1 != nil || cur2 != nil {
			ans.Next = new(ListNode)
			ans = ans.Next
		}
	}
	for cur1 != nil {
		ans.Val = (cur1.Val + c) % 10
		c = (cur1.Val + c) / 10
		cur1 = cur1.Next
		if cur1 != nil {
			ans.Next = new(ListNode)
			ans = ans.Next
		}
	}
	for cur2 != nil {
		ans.Val = (cur2.Val + c) % 10
		c = (cur2.Val + c) / 10
		cur2 = cur2.Next
		if cur2 != nil {
			ans.Next = new(ListNode)
			ans = ans.Next
		}
	}
	if c != 0 {
		ans.Next = new(ListNode)
		ans = ans.Next
		ans.Val = c
	}
	return res
}

type BSTIterator struct {
	arr []int
	idx int
}

func ConstructorBST(root *TreeNode) BSTIterator {
	temp := make([]int, 0)
	var recurence func(root *TreeNode)
	recurence = func(root *TreeNode) {
		if root == nil {
			return
		}
		recurence(root.Left)
		temp = append(temp, root.Val)
		recurence(root.Right)
	}
	recurence(root)
	return BSTIterator{temp, 0}
}

func (this *BSTIterator) Next() int {
	ans := this.arr[this.idx]
	this.idx++
	return ans
}

func (this *BSTIterator) HasNext() bool {
	return len(this.arr) > this.idx
}

type SeatManager struct {
	heap digitHeap
}

func ConstructorSeat(n int) SeatManager {
	h := digitHeap{}
	heap.Init(&h)
	for i := 0; i < n; i++ {
		heap.Push(&h, i)
	}
	return SeatManager{h}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(&this.heap).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&this.heap, seatNumber)
}
