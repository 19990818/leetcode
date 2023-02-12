package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func pickGifts(gifts []int, k int) int64 {
	h1 := &digitHeap{}
	heap.Init(h1)
	for _, v := range gifts {
		heap.Push(h1, v)
	}
	for i := 0; i < k; i++ {
		t := heap.Pop(h1).(int)
		heap.Push(h1, int(math.Sqrt(float64(t))))
	}
	res := 0
	for h1.Len() > 0 {
		res += heap.Pop(h1).(int)
	}
	return int64(res)
}

type digitHeap []int

func (h digitHeap) Len() int {
	return len(h)
}
func (h digitHeap) Less(i, j int) bool {
	return h[i] > h[j]
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

func vowelStrings(words []string, queries [][]int) []int {
	sum := make([]int, len(words)+1)
	for i, v := range words {
		sum[i+1] = sum[i]
		if startAndEndVowel(v) {
			sum[i+1] += 1
		}
	}
	// fmt.Println(sum)
	res := make([]int, len(queries))
	for i, v := range queries {
		res[i] = sum[v[1]+1] - sum[v[0]]
	}
	return res
}
func startAndEndVowel(a string) bool {
	m := make(map[rune]bool)
	for _, v := range "aeiou" {
		m[v] = true
	}
	return m[rune(a[0])] && m[rune(a[len(a)-1])]
}

func minCapability(nums []int, k int) int {
	left, right := 0, int(1e9+1)
	for left < right {
		mid := (left + right) >> 1
		if judge(mid, nums, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	fmt.Println(left)
	sort.Ints(nums)
	for _, v := range nums {
		if v >= left {
			return v
		}
	}
	return -1
}
func judge(mid int, nums []int, k int) bool {
	i := 0
	cnt := 0
	for i < len(nums) {
		if nums[i] <= mid {
			//选择这个
			i += 2
			cnt++
		} else {
			i++
		}
	}
	return cnt >= k
}

func minCost(basket1 []int, basket2 []int) int64 {
	// sort.Ints(basket1)
	// sort.Ints(basket2)
	diff := make(map[int]int)
	mi := basket1[0]
	for _, v := range basket1 {
		diff[v]++
		mi = min(mi, v)
	}
	for _, v := range basket2 {
		diff[v]--
		mi = min(mi, v)
	}
	pairs := 0
	total := make([]int, 0)
	for k, v := range diff {
		if v%2 == 1 {
			return -1
		}
		if v > 0 {
			pairs += v / 2
		}
		if v != 0 {
			total = append(total, k)
		}

	}
	sort.Ints(total)
	fmt.Println(total, diff, pairs)
	res := 0
	i := 0
	for pairs > 0 && i < len(total) {
		if diff[total[i]] != 0 {
			res += min(total[i], 2*mi) * abs(diff[total[i]]) / 2
			pairs -= abs(diff[total[i]]) / 2
		}
		i++
	}
	return int64(res)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

[4,2,2,2]
[1,4,1,2]
[2,3,4,1]
[3,2,5,1]
[4,4,4,4,3]
[5,5,5,5,3]
[84,80,43,8,80,88,43,14,100,88]
[32,32,42,68,68,100,42,84,14,8]
