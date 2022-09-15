package main

import (
	"container/heap"
)

// #855 考场就坐
//此题实际上可以将区间段作为元素加入堆中 然后以距离作为优先级
//每次取出最大长度的区间段 其中两个边界需要特殊处理 当存在边界的区间段时
//他们的距离不是取中点值，而是直接这段区间的距离。为了较好处理各区间段
//之间的边界问题，采用左闭右开区间
//因为是半开半闭区间 计算中间值需减去1，计算距离需加上1，距离需加上1是因为
//实际上左边可以视为(left-1,righ）
type ExamRoom struct {
	h     *h
	pairs map[int]int
	len   int
}

func ConstructorSeat(n int) ExamRoom {
	temp := &h{}
	heap.Init(temp)
	heap.Push(temp, interval{0, n, n})
	pairs := make(map[int]int)
	pairs[0] = n
	pairs[n] = 0
	return ExamRoom{
		temp,
		pairs,
		n,
	}
}
func (this *ExamRoom) add(l, r int) {
	var distance int
	if l == 0 || r == this.len {
		distance = r - l
	} else {
		distance = (r - l + 1) >> 1
	}
	heap.Push(this.h, interval{l, r, distance})
	this.pairs[l] = r
	this.pairs[r] = l
}
func (this *ExamRoom) Seat() int {
	var in interval
	for this.h.Len() > 0 {
		in = heap.Pop(this.h).(interval)
		if this.pairs[in.left] == in.right && this.pairs[in.right] == in.left {
			break
		}
	}
	var p int
	//fmt.Println(in)
	if in.left == 0 {
		p = 0
	} else if in.right == this.len {
		p = in.right - 1
	} else {
		p = (in.right + in.left - 1) >> 1
	}
	this.add(in.left, p)
	this.add(p+1, in.right)
	return p
}

func (this *ExamRoom) Leave(p int) {
	left, right := this.pairs[p], this.pairs[p+1]
	this.add(left, right)
}

type interval struct {
	left     int
	right    int
	distance int
}
type h []interval

func (h h) Len() int {
	return len(h)
}
func (h h) Less(i, j int) bool {
	if h[i].distance > h[j].distance {
		return true
	}
	if h[i].distance == h[j].distance {
		return h[i].left < h[j].left
	}
	return false
}
func (h *h) Pop() interface{} {
	old := *h
	temp := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return temp
}
func (h *h) Push(x interface{}) {
	*h = append(*h, x.(interval))
}
func (h h) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
