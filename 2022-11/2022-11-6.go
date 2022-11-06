package main

import (
	"container/heap"
	"math"
	"sort"
)

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i], nums[i+1] = nums[i]*2, 0
		}
	}
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			res = append(res, nums[i])
		}
	}
	for i := len(res); i < len(nums); i++ {
		res = append(res, 0)
	}
	return res
}

func maximumSubarraySum(nums []int, k int) int64 {
	// m 维持区间内出现次数
	m := make(map[int]int)
	ans, sum := int64(0), int64(0)
	for i := 0; i < k; i++ {
		m[nums[i]]++
		sum += int64(nums[i])
	}
	if len(m) == k {
		ans = max64(ans, sum)
	}
	for i := 0; i < len(nums)-k; i++ {
		m[nums[i]]--
		if m[nums[i]] == 0 {
			delete(m, nums[i])
		}
		sum = sum - int64(nums[i]) + int64(nums[i+k])
		m[nums[i+k]]++
		if len(m) == k {
			ans = max64(ans, sum)
		}
	}
	return ans
}
func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func totalCost(costs []int, k int, candidates int) int64 {
	if 2*candidates >= len(costs) {
		sort.Ints(costs)
		ans := int64(0)
		for i := 0; i < k; i++ {
			ans += int64(costs[i])
		}
		return ans
	}
	h1, h2 := digitHeap{}, digitHeap{}
	heap.Init(&h1)
	heap.Init(&h2)
	for i := 0; i < candidates; i++ {
		heap.Push(&h1, costs[i])
		heap.Push(&h2, costs[len(costs)-1-i])
	}
	ans := int64(0)
	i, j := candidates-1, len(costs)-candidates
	for k > 0 {
		temp1, temp2 := math.MaxInt64, math.MaxInt64
		if len(h1) > 0 {
			temp1 = heap.Pop(&h1).(int)
		}
		if len(h2) > 0 {
			temp2 = heap.Pop(&h2).(int)
		}
		if temp1 <= temp2 {
			ans += int64(temp1)
			i++
			if i < j {
				heap.Push(&h1, costs[i])
			}
			heap.Push(&h2, temp2)
		} else {
			ans += int64(temp2)
			j--
			if i < j {
				heap.Push(&h2, costs[j])
			}
			heap.Push(&h1, temp1)
		}
		k--
	}
	return ans
}

//构造小根堆函数
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

// 我们首先需要将机器人和工厂进行一个排序，然后每个工厂都会将要修理一个连续区间的机器人
// 因为如果不连续的话，我们交换将会得到一个更大的值
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})
	m, n := len(robot), len(factory)
	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, m+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}
	//fmt.Println(dp)
	for i := 0; i < n; i++ {
		dp[i][m] = 0
	}
	for j := 0; j < m; j++ {
		for k := 1; k <= factory[n-1][1]; k++ {
			dp[n-1][m-k] = int64(abs(robot[m-k]-factory[n-1][0])) + dp[n-1][m-k+1]
		}
	}
	//fmt.Println(dp)
	for i := n - 2; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			sum := int64(0)
			dp[i][j] = min64(dp[i][j], dp[i+1][j])
			for k := 1; k <= factory[i][1] && j+k <= m; k++ {
				sum += int64(abs(robot[j+k-1] - factory[i][0]))
				if dp[i+1][j+k] != math.MaxInt64 {
					dp[i][j] = min64(dp[i][j], dp[i+1][j+k]+sum)
				}
			}
		}
	}
	//fmt.Println(dp)
	return dp[0][0]
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
