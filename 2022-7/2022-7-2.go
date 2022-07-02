package main

import (
	"container/heap"
)

type digit []int

func (h digit) Len() int {
	return len(h)
}
func (h digit) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h digit) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *digit) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *digit) Pop() interface{} {
	old := *h
	num := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return num
}
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	maxDis := startFuel
	stack := digit{}
	heap.Init(&stack)
	stations = append(stations, []int{target, 0})
	ans := 0
	for _, val := range stations {
		if val[0] >= maxDis {
			for maxDis < val[0] && len(stack) > 0 {
				ans++
				maxNum := heap.Pop(&stack).(int)
				//fmt.Println(maxNum)
				maxDis += maxNum
			}
			if maxDis < val[0] {
				return -1
			}
		}
		heap.Push(&stack, val[1])
	}
	return ans
}

func calculateMinimumHP(dungeon [][]int) int {
	//当我们正向选择的时候会遇到一个问题
	//我们无法将所有正值进行保存，然后导致子问题和父问题
	//出现割裂的情况 比如
	//-2 -8 10
	//-5 -2 3
	//-3 0 -9在这种情况下我们每次保存相邻的值-9上方最小解为9
	//左边最优解也为9 那么-9处最优解为18 而-2-8 10 3 -9最优解为10
	//出现这样的情况是我们没有对正值进行一个合理的利用 在此处10 3失去了消掉-9的机会
	//当我们选择反向的时候，遇到前面的正值可以充分消除后面的负值
	//当可以进行消除的时候 相当于去掉了这种割裂
	m := len(dungeon)
	n := len(dungeon[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			temp := max(dp[i+1][j], dp[i][j+1])
			if i+1 == m && j+1 != n {
				temp = dp[i][j+1]
			}
			if i+1 != m && j+1 == n {
				temp = dp[i+1][j]
			}
			if dungeon[i][j] < 0 {
				dp[i][j] = min(temp, 0) + dungeon[i][j]
			} else {
				dp[i][j] = temp + dungeon[i][j]
			}
		}
	}
	//fmt.Println(dp)
	if dp[0][0] >= 0 {
		return 1
	}
	return -dp[0][0] + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(k int, prices []int) int {
	if len(prices) == 1 || len(prices) == 0 || k < 1 {
		return 0
	}
	profits := make([]int, len(prices)-1)
	for i := 0; i < len(prices)-1; i++ {
		profits[i] = prices[i+1] - prices[i]
	}
	n := len(profits)
	//动态规划 dp表示必须选以j为结束的最大利润
	//dp2表示不选以j为结束的最大利润
	//必须选和不选
	dp := make([][]int, k)
	dp2 := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp2[i] = make([]int, n+1)
	}

	for j := n - 1; j >= 0; j-- {
		for i := 0; i < k; i++ {
			if i == 0 || j == n-1 {
				//当只能进行一次买卖 或者是最开始的时候
				//必须选择j元素，其最大值为只选该值，和前一个元素的最大值加上该值
				dp[i][j] = max(profits[j]+dp[i][j+1], profits[j])
			} else {
				//可以进行多次买卖
				//就会增加一个多进行一次操作 多进行的一次操作取目前的值加上距离为2的交易
				//能得到的最大值
				dp[i][j] = max(profits[j]+dp[i][j+1], profits[j])
				dp[i][j] = max(dp[i][j], profits[j]+dp2[i-1][j+2])
			}
			//非必须的情况下，加上不选的情况
			dp2[i][j] = max(dp[i][j], dp2[i][j+1])
		}
	}
	return max(dp2[k-1][0], 0)
}
