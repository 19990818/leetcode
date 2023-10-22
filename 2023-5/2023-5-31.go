package main

import (
	"container/heap"
	"math"
)

func mctFromLeafValues2(arr []int) int {
	t := h{}
	heap.Init(&t)
	for _, v := range arr {
		heap.Push(&t, v)
	}
	for len(t) > 1 {
		t1 := heap.Pop(&t).(int)
		t2 := heap.Pop(&t).(int)
		heap.Push(&t, t1*t2)
	}
	return heap.Pop(&t).(int)
}

type h []int

func (t h) Len() int {
	return len(t)
}
func (t h) Less(i, j int) bool {
	return t[i] < t[j]
}
func (t h) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t *h) Push(x interface{}) {
	*t = append(*t, x.(int))
}
func (t *h) Pop() interface{} {
	old := *t
	res := old[0]
	*t = old[1:]
	return res
}

func mctFromLeafValues(arr []int) int {
	dpS := make([][]int, len(arr))
	dpM := make([][]int, len(arr))
	for i := range dpS {
		dpS[i] = make([]int, len(arr)+1)
		dpM[i] = make([]int, len(arr)+1)
		dpM[i][i+1] = arr[i]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		for j := i + 1; j <= len(arr); j++ {
			dpS[i][j] = 0
			for k := i + 1; k < j; k++ {
				if dpS[i][j] == 0 {
					dpS[i][j] = dpS[i][k] + dpM[i][k]*dpM[k][j] + dpS[k][j]
				}
				dpS[i][j] = min(dpS[i][k]+dpM[i][k]*dpM[k][j]+dpS[k][j], dpS[i][j])
				//fmt.Println(dpS[i][k]+dpM[i][k]*dpM[k][j]+dpS[k][j],dpS[i][j])
			}
			dpM[i][j] = max(dpM[i][j-1], arr[j-1])
		}
	}
	return dpS[0][len(arr)]
}
func dfs(arr []int) (sum int, mnum int) {
	if len(arr) == 0 {
		return 0, 0
	}
	if len(arr) == 1 {
		return sum, arr[0]
	}
	sum = math.MaxInt64
	for i := 1; i < len(arr); i++ {
		ls, lm := dfs(arr[0:i])
		rs, rm := dfs(arr[i:])
		if sum > ls+rs+lm*rm {
			sum = ls + rs + lm*rm
			mnum = max(lm, rm)
		}
	}
	return sum, mnum
}
