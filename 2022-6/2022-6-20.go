package main

import "github.com/emirpasic/gods/trees/redblacktree"

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][n] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	//fmt.Println(dp)
	return dp[0][0]
}

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	profits := make([]int, len(prices)-1)
	for i := 0; i < len(prices)-1; i++ {
		profits[i] = prices[i+1] - prices[i]
	}
	n := len(profits)

	count := 2
	//必须选和不选
	dp := make([][]int, count)
	dp2 := make([][]int, count)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp2[i] = make([]int, n+1)
	}
	dp[0][n] = 0
	dp[1][n] = 0
	dp2[0][n] = 0
	dp2[0][n] = 0
	for j := n - 1; j >= 0; j-- {
		for i := 0; i < count; i++ {
			if i == 0 || j == n-1 {
				dp[i][j] = max(profits[j]+dp[i][j+1], profits[j])
				dp2[i][j] = max(dp2[i][j], profits[j])
				dp2[i][j] = max(profits[j]+dp[i][j+1], dp2[i][j])
			} else {
				dp[i][j] = max(profits[j]+dp[i][j+1], profits[j])
				dp[i][j] = max(dp[i][j], profits[j]+dp2[i-1][j+2])
				dp2[i][j] = max(dp2[i][j], profits[j])
				dp2[i][j] = max(profits[j]+dp[i][j+1], dp2[i][j])
				dp2[i][j] = max(dp2[i][j], profits[j]+dp2[i-1][j+2])
			}
			dp2[i][j] = max(dp2[i][j], dp2[i][j+1])
		}
	}
	return max(dp2[1][0], 0)
}

type RangeModule struct {
	*redblacktree.Tree
}

func Constructor2() RangeModule {
	return RangeModule{redblacktree.NewWithIntComparator()}
}

func (this *RangeModule) AddRange(left int, right int) {
	if node, ok := this.Floor(left); ok {
		r := node.Value.(int)
		//在一个节点内部的情况
		if right <= r {
			return
		}
		//开始节点在一个节点内部的情况
		if r >= left {
			left = node.Key.(int)
			this.Remove(node.Key)
		}
	}
	for node, ok := this.Ceiling(left); ok && right >= node.Key.(int); node, ok = this.Ceiling(left) {
		right = max(right, node.Value.(int))
		this.Remove(node.Key)
	}
	this.Put(left, right)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	node, ok := this.Floor(left)
	if ok && node.Value.(int) >= right {
		return true
	}
	return false
}

func (this *RangeModule) RemoveRange(left int, right int) {
	//首先讨论在一个节点内部的情况
	if node, ok := this.Floor(left); ok {
		l, r := node.Key.(int), node.Value.(int)
		if right <= r {
			if left == l {
				this.Remove(node.Key)
			} else {
				node.Value = left
			}
			if right != r {
				this.Put(right, r)
			}
			return
		}
		if r > left {
			node.Value = left
		}
	}
	for node, ok := this.Ceiling(left); ok && right > node.Key.(int); node, ok = this.Ceiling(left) {
		r := node.Value.(int)
		this.Remove(node.Key)
		if right < r {
			this.Put(right, r)
			break
		}
	}
}
