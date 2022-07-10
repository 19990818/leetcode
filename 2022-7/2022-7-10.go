package main

import (
	"container/heap"
	"math"
	"sort"
	"strings"
)

func fillCups(amount []int) int {
	//一共就三个杯子 最大的搞定
	sort.Ints(amount)
	ans := 0
	ans += amount[2]
	if amount[0]+amount[1] <= amount[2] {
		return ans
	}
	for amount[2] > 0 {
		if amount[0] < amount[1] {
			amount[1]--
		} else {
			amount[0]--
		}
		amount[2]--
	}

	return ans + max(amount[0], amount[1])
}

type SmallestInfiniteSet struct {
	m map[int]int
	h digit
}

func Constructor() SmallestInfiniteSet {
	m := make(map[int]int)
	h := digit{}
	heap.Init(&h)
	for i := 1; i <= 1000; i++ {
		m[i] = 1
		heap.Push(&h, i)
	}
	return SmallestInfiniteSet{m, h}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	cur := heap.Pop(&this.h).(int)
	delete(this.m, cur)
	return cur
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.m[num] == 0 {
		this.m[num] = 1
		heap.Push(&this.h, num)
	}
}

func canChange(start string, target string) bool {
	//统计每个lr左边和右边的空格
	//l右边的空格应该>=start r左边的空格应该>=start
	var t1, t2 strings.Builder
	for i := range start {
		if start[i] != '_' {
			t1.WriteByte(start[i])
		}
		if target[i] != '_' {
			t2.WriteByte(target[i])
		}
	}
	if t1.String() != t2.String() {
		return false
	}
	cnt, cnt2 := 0, 0
	RLSpace, tarRLSpace := make([]int, 0), make([]int, 0)
	LRSpace, tarLRSpace := make([]int, 0), make([]int, 0)
	for i := range start {
		if start[i] == '_' {
			cnt++
		}
		if target[i] == '_' {
			cnt2++
		}
		if start[i] == 'R' {
			RLSpace = append(RLSpace, cnt)
		}
		if target[i] == 'R' {
			tarRLSpace = append(tarRLSpace, cnt2)
		}
	}
	cnt, cnt2 = 0, 0
	for i := len(start) - 1; i >= 0; i-- {
		if start[i] == '_' {
			cnt++
		}
		if target[i] == '_' {
			cnt2++
		}
		if start[i] == 'L' {
			LRSpace = append(LRSpace, cnt)
		}
		if target[i] == 'L' {
			tarLRSpace = append(tarLRSpace, cnt2)
		}
	}
	//fmt.Println(tarLRSpace,LRSpace)
	//fmt.Println(tarRLSpace)
	if len(tarLRSpace) != len(LRSpace) || len(tarRLSpace) != len(tarRLSpace) {
		return false
	}
	for i := range tarLRSpace {
		if tarLRSpace[i] < LRSpace[i] {
			return false
		}
	}
	for i := range tarRLSpace {
		if tarRLSpace[i] < RLSpace[i] {
			return false
		}
	}
	return true
}

func cherryPickup(grid [][]int) int {
	//首先考虑直接两次dp 每次得到最大可能值
	//但是局部最大无法保证全局最大
	//然后提示将两个人都视为从00出发 得到每个点所能对应的最大值
	//状态转移方程存在问题，按照次数进行处理会存在重复
	//因为每个人实际上最多移动的距离为2n-2，可以以此处理
	//将每个人移动相同距离得到能够得到最大值作为origin问题
	//移动距离-1可以作为子问题，xy值表示坐标 距离固定可以不需要y
	//三维动态规划 1 距离 2 a位置 3 b位置
	//实际上可以抽象成一个闭合的圆，ab位置等效 可以假设xa<=xb
	n := len(grid)
	dp := make([][][]int, 2*n-1)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MinInt64
			}
		}
	}
	dp[0][0][0] = grid[0][0]
	for dis := 1; dis < 2*n-1; dis++ {
		for x1 := max(0, dis-n+1); x1 < min(n, dis+1); x1++ {
			y1 := dis - x1
			if grid[x1][y1] == -1 {
				continue
			}
			for x2 := x1; x2 < min(n, dis+1); x2++ {
				y2 := dis - x2
				if grid[x2][y2] == -1 {
					continue
				}
				//两个都向右移动x不变
				res := dp[dis-1][x1][x2]
				//x1向右移动 x2向下移动
				if x2 > 0 {
					res = max(res, dp[dis-1][x1][x2-1])
				}
				//x1向下移动 x2向右移动
				if x1 > 0 {
					res = max(res, dp[dis-1][x1-1][x2])
				}
				//都向下移动
				if x1 > 0 && x2 > 0 {
					res = max(res, dp[dis-1][x1-1][x2-1])
				}
				dp[dis][x1][x2] = res + grid[x1][y1]
				if x1 != x2 {
					dp[dis][x1][x2] += grid[x2][y2]
				}
			}
		}
	}
	return max(0, dp[2*n-2][n-1][n-1])
}
