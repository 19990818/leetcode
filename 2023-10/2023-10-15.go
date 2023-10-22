package main

import "container/heap"

type st []int

func maxKelements(nums []int, k int) int64 {
	h := st{}
	heap.Init(&h)
	for _, v := range nums {
		heap.Push(&h, v)
	}
	res := 0
	for i := 0; i < k; i++ {
		cur := heap.Pop(&h).(int)
		res += cur
		heap.Push(&h, (cur+2)/3)
	}
	return int64(res)
}
func (s st) Len() int {
	return len(s)
}
func (s st) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s st) Less(i, j int) bool {
	return s[i] > s[j]
}
func (s *st) Pop() interface{} {
	old := *s
	res := old[len(old)-1]
	old = old[0 : len(old)-1]
	*s = old
	return res
}
func (s *st) Push(x interface{}) {
	old := *s
	old = append(old, x.(int))
	*s = old
}
