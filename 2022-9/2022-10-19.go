package main

import (
	"sort"
)

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	//实际上就是找numwanted个数字 满足uselimit要求
	type pair struct {
		val   int
		label int
	}
	pairs := make([]pair, len(values))
	for i := 0; i < len(values); i++ {
		pairs[i] = pair{values[i], labels[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})
	cnt := make(map[int]int)
	ans := 0
	for i := len(pairs) - 1; i >= 0 && numWanted > 0; i-- {
		if cnt[pairs[i].val] < useLimit {
			cnt[pairs[i].val]++
			numWanted--
			ans += pairs[i].val
		}
	}
	return ans
}

func shortestPathBinaryMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[m-1][n-1] == 1 {
		return -1
	}
	travel := make(map[int]int)
	tos := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	travel[0] = 1
	cnt := 0
	queue := []int{0}
	for {
		temp := make([]int, 0)
		cnt++
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur == m*n-1 {
				return cnt
			}
			x, y := cur/n, cur%n
			for _, to := range tos {
				if x+to[0] < m && x+to[0] >= 0 &&
					y+to[1] < n && y+to[1] >= 0 &&
					travel[(x+to[0])*n+y+to[1]] == 0 && grid[x+to[0]][y+to[1]] == 0 {
					travel[(x+to[0])*n+y+to[1]] = 1
					temp = append(temp, (x+to[0])*n+y+to[1])
				}
			}
		}
		if len(temp) == 0 {
			break
		}
		queue = temp
	}
	return -1
}

func pathInZigZagTree(label int) []int {
	temp := make([]int, 0)
	temp = append(temp, label)
	f := func(lable int) int {
		for i := 31; i >= 0; i-- {
			if lable >= 1<<i {
				return 1 << i
			}
		}
		return 0
	}
	for label > 1 {
		label /= 2
		sum := f(label)*3 - 1
		temp = append(temp, sum-label)
		label = sum - label
	}
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	return temp
}

type StockSpanner struct {
	cur   int
	stack []int
	num   []int
}

func Constructor2() StockSpanner {
	return StockSpanner{0, make([]int, 0), make([]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	res := 0
	for len(this.stack) > 0 && this.num[len(this.num)-1] <= price {
		this.num = this.num[0 : len(this.num)-1]
		this.stack = this.stack[0 : len(this.stack)-1]
	}
	if len(this.stack) == 0 {
		res = this.cur + 1
	} else {
		res = this.cur - this.stack[len(this.stack)-1]
	}
	this.stack = append(this.stack, this.cur)
	this.num = append(this.num, price)
	this.cur++
	return res
}
