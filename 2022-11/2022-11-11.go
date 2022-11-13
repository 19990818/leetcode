package main

import (
	"container/heap"
)

func eatenApples(apples []int, days []int) int {
	res := 0
	m := make(map[int]int)
	h := expired{}
	heap.Init(&h)
	for i := range apples {
		if m[i+days[i]] == 0 {
			heap.Push(&h, i+days[i])
		}
		m[i+days[i]] += apples[i]
		res = eat(&h, res, m, i)
	}
	for i := len(apples); len(h) > 0; i++ {
		res = eat(&h, res, m, i)
	}
	return res
}
func eat(h *expired, res int, m map[int]int, d int) int {
	temp := heap.Pop(h).(int)
	for temp <= d && len(*h) > 0 {
		temp = heap.Pop(h).(int)
	}
	m[temp]--
	if m[temp] >= 0 && temp > d {
		res++
		if m[temp] > 0 {
			heap.Push(h, temp)
		}
	}
	return res
}

type expired []int

func (e expired) Len() int {
	return len(e)
}
func (e expired) Less(i, j int) bool {
	return e[i] < e[j]
}
func (e expired) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e *expired) Pop() interface{} {
	old := *e
	temp := old[len(old)-1]
	*e = old[0 : len(old)-1]
	return temp
}
func (e *expired) Push(x interface{}) {
	*e = append(*e, x.(int))
}
