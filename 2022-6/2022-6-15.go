package main

import "sort"

func minDistance(word1 string, word2 string) int {
	//实际上是看word1 对于word2的最长子串
	// if len(word1) == 0 || len(word2) == 0 {
	// 	return abs(len(word1) - len(word2))
	// }
	// if word1[0] == word2[0] {
	// 	return minDistance(word1[1:], word2[1:])
	// }
	//在此处首先只考虑删除和替换 未考虑插入
	// return min(minDistance(word1[1:], word2)+1, minDistance(word1[1:], word2[1:])+1)
	//插入情况为word,word2[1:]
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n2; j++ {
		dp[0][j] = j
	}
	//以i,j分别结束的结果
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			}
		}
	}
	//fmt.Println(dp)
	return dp[n1][n2]
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	colMax, rowMax := make([]int, len(grid[0])), make([]int, len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			rowMax[i] = max(rowMax[i], grid[i][j])
			colMax[j] = max(colMax[j], grid[i][j])
		}
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			ans += min(colMax[j], rowMax[i]) - grid[i][j]
		}
	}
	return ans
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	var lessAndEqualX func(nums []int, x int) int
	lessAndEqualX = func(nums []int, x int) int {
		ans := 0
		i, j := 0, 1
		for ; j < len(nums); j++ {
			for nums[j]-nums[i] > x {
				i++
			}
			ans += j - i
		}
		return ans
	}
	left, right := 0, int(1e6)
	var mid int
	for left < right {
		mid = (right-left)>>2 + left
		//fmt.Println(nums,mid,lessAndEqualX(nums, mid))
		if lessAndEqualX(nums, mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
