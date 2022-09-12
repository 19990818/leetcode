package main

import (
	"container/heap"
	"sort"
)

func mostFrequentEven(nums []int) int {
	m := make(map[int]int)
	ans := -1
	maxCnt := 0
	sort.Ints(nums)
	for _, val := range nums {
		if val%2 == 0 {
			m[val]++
			if m[val] > maxCnt {
				ans = val
				maxCnt = m[val]
			}
		}
	}
	return ans
}

func partitionString(s string) int {
	ans := 1
	m := make(map[rune]int)
	for _, val := range s {
		if m[val] != 0 {
			m = make(map[rune]int)
			ans++
		}
		m[val] = 1
		//fmt.Println(m)
	}
	return ans
}

func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//fmt.Println(intervals)
	intervalheap := &digitHeap{}
	heap.Init(intervalheap)
	heap.Push(intervalheap, intervals[0][1])
	for i := 1; i < len(intervals); i++ {
		temp := heap.Pop(intervalheap).(int)
		//fmt.Println(intervalheap,temp)
		//有重叠
		if intervals[i][0] <= temp {
			heap.Push(intervalheap, temp)
		}
		heap.Push(intervalheap, intervals[i][1])
	}
	return len(*intervalheap)
}

type digitHeap []int

func (h digitHeap) Len() int {
	return len(h)
}
func (h digitHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h digitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *digitHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *digitHeap) Pop() interface{} {
	old := *h
	num := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return num
}
