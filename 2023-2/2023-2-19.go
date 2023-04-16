package main

import (
	"container/heap"
	"math"
	"sort"
)

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	var h1 h
	heap.Init(&h1)
	for _, v := range classes {
		heap.Push(&h1, v)
	}
	for i := 0; i < extraStudents; i++ {
		t := heap.Pop(&h1)
		cur := []int{t.([]int)[0] + 1, t.([]int)[1] + 1}
		heap.Push(&h1, cur)
	}
	res := float64(0)
	for len(h1) > 0 {
		t := heap.Pop(&h1)
		res += float64(t.([]int)[0]) / float64(t.([]int)[1])
	}
	return res
}

type h [][]int

func (t h) Len() int {
	return len(t)
}
func (t h) Less(i, j int) bool {
	return t[j][1]*(t[i][1]-t[i][0])*(t[j][1]+1) > t[i][1]*(t[i][1]+1)*(t[j][1]-t[j][0])
}
func (t h) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t *h) Pop() interface{} {
	old := *t
	res := old[len(old)-1]
	old = old[0 : len(old)-1]
	*t = old
	return res
}
func (t *h) Push(x interface{}) {
	*t = append(*t, x.([]int))
}

func leftRigthDifference(nums []int) []int {
	leftSum := make([]int, len(nums)+1)
	sum := 0
	for i, v := range nums {
		sum += v
		leftSum[i+1] = leftSum[i] + v
	}
	res := make([]int, len(nums))
	for i, _ := range nums {
		res[i] = abs(sum - leftSum[i+1] - leftSum[i])
	}
	return res
}

func divisibilityArray(word string, m int) []int {
	res := make([]int, len(word))
	c := 0
	for i, v := range word {
		c = (c*10 + int(v-'0')) % m
		if c == 0 {
			res[i] = 1
		}
	}
	return res
}

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	right, left := len(nums)-1, len(nums)/2-1
	res := 0
	for ; right > len(nums)/2-1 && left >= 0; right-- {
		for left >= 0 && nums[right] < nums[left]*2 {
			left--
		}
		if left >= 0 {
			res += 2
			left--
		}
	}
	return res
}

func minimumTime(grid [][]int) int {
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}
	tos := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	travel := make(map[int]int)
	m, n := len(grid), len(grid[0])
	startTime := math.MaxInt64
	endTime := sort.Search(1e5+m+n, func(ti int) bool {
		if ti < grid[m-1][n-1] || ti < m+n-2 {
			return false
		}
		q := []int{m*n - 1}
		travel[m*n-1] = ti
		for cur := ti - 1; len(q) > 0; cur-- {
			t := q
			q = nil
			for _, v := range t {
				x, y := v/n, v%n
				for _, to := range tos {
					xn, yn := x+to[0], y+to[1]
					if !(xn < m && xn >= 0 && yn < n && yn >= 0) {
						continue
					}

					if travel[xn*n+yn] != ti && grid[xn][yn] <= cur {
						if xn == 0 && yn == 0 {
							startTime = min(startTime, cur)
							return true
						}
						q = append(q, xn*n+yn)
						travel[xn*n+yn] = ti
					}
				}
			}
		}
		return false
	})
	return endTime + startTime%2
}
