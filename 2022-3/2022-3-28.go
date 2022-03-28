package main

import "sort"

//动态规划 难点是将各种数字的选择情况以状态的形式表示
//这样可以使用位运算很简单的完成是否已经尝试过
//如果当前状态存在一个选择可以使值直接大于想要的 直接获胜
//或者是当前状态选择之后的状态是一种必输的状态 那么也可以获胜
//这样就可以完成对数组状态的存储
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	var dfs func(state int, desiredTotal int, dp []int) bool
	dfs = func(state, desiredTotal int, dp []int) bool {
		//fmt.Println(dp)
		if dp[state] != 0 {
			if dp[state] == 1 {
				return true
			}
			return false
		}
		for i := 1; i <= maxChoosableInteger; i++ {
			cur := 1 << (i - 1)
			if cur&state != 0 {
				continue
			}
			if i >= desiredTotal || !dfs(state|cur, desiredTotal-i, dp) {
				dp[state] = 1
				//fmt.Println(i)
				return true
			}
		}
		dp[state] = 2
		return false
	}
	dp := make([]int, 1<<maxChoosableInteger)
	return dfs(0, desiredTotal, dp)
}

//处理子串 如果是从前到后 那么就需要将以当前位置为结尾作为一个参考
//当以这个结尾的长度为多少 也就是意味着这样的子串有多少 动态规划
func findSubstringInWraproundString(p string) int {
	dp := make([]int, 26)
	ans := 0
	maxLetter := make([]int, 26)
	dp[p[0]-'a'] = 1
	maxLetter[p[0]-'a'] = 1
	for i := 1; i < len(p); i++ {
		if p[i]-p[i-1] == 1 || (p[i] == 'a' && p[i-1] == 'z') {
			dp[p[i]-'a'] = dp[p[i-1]-'a'] + 1
		} else {
			dp[p[i]-'a'] = 1
		}
		maxLetter[p[i]-'a'] = max(maxLetter[p[i]-'a'], dp[p[i]-'a'])
	}
	//fmt.Println(maxLetter)
	for _, val := range maxLetter {
		ans += val
	}
	return ans
}

//dfs 从大的开始处理 可以减少时间 四个直接一起处理 不要分开处理
//每个元素在各组中出现概率相同
func makesquare(matchsticks []int) bool {
	sum := 0
	for _, val := range matchsticks {
		sum += val
	}
	if sum%4 != 0 {
		return false
	}
	target := sum / 4
	a := make([]int, 4)
	sort.Ints(matchsticks)
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == -1 {
			return a[0] == a[1] && a[1] == a[2] && a[2] == a[3]
		}
		for i := 0; i < 4; i++ {
			if a[i]+matchsticks[index] <= target {
				a[i] += matchsticks[index]
				if dfs(index - 1) {
					return true
				}
				a[i] -= matchsticks[index]
			}
		}
		return false
	}
	return dfs(len(matchsticks) - 1)
}
