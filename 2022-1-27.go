package main

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	for i := n; i >= 0; i-- {
		if i == len(citations) {
			if citations[n-i] >= i {
				return i
			}
		} else if i == 0 {
			if citations[n-i-1] <= i {
				return i
			}
		} else {
			if i >= citations[n-i-1] && i <= citations[n-i] {
				return i
			}
		}
	}
	return -1
}

func hIndex2(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	left, right := n, 0
	for left >= right {
		i := (left-right)/2 + right
		if i == len(citations) {
			if citations[n-i] >= i {
				return i
			}
			left = i - 1
		} else if i == 0 {
			if citations[n-i-1] <= i {
				return i
			}
			right = i + 1
		} else {
			if i >= citations[n-i-1] && i <= citations[n-i] {
				return i
			} else if i < citations[n-i-1] {
				right = i + 1
			} else {
				left = i - 1
			}
		}
	}
	return -1
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if dp[i] == 0 {
				dp[i] = dp[i-j*j] + 1
			} else {
				dp[i] = min(dp[i], dp[i-j*j]+1)
			}
		}
	}
	return dp[n]
}

func findDuplicate(nums []int) int {
	// 通过构建边得到像链表一样的结构 然后就成为了找到环的入口
	// 只有一个超过两个的入度作为环的节点
	slow, fast := 0, 0
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}
	return num%10 != 0
}
