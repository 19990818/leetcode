package main

import "sort"

//深度优先遍历和动态规划组合
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][]int, m)
	mod := int(1e9 + 7)
	ans := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[startRow][startColumn] = 1
	for maxMove > 0 {
		cur := make([][]int, m)
		for i := 0; i < m; i++ {
			cur[i] = make([]int, n)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				near := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
				for _, val := range near {
					if i+val[0] < m && i+val[0] >= 0 && j+val[1] >= 0 && j+val[1] < n {
						cur[i+val[0]][j+val[1]] = (cur[i+val[0]][j+val[1]] + dp[i][j]) % mod
					} else {
						ans = (ans + dp[i][j]) % mod
					}
				}
			}
		}
		maxMove--
		dp = cur
		//fmt.Println(dp)
	}
	return ans
}

func findUnsortedSubarray(nums []int) int {
	original := append([]int{}, nums...)
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < len(nums) && nums[left] == original[left] {
		left++
	}
	for right >= 0 && nums[right] == original[right] {
		right--
	}
	if right <= left {
		return 0
	}
	return right - left + 1
}

func findUnsortedSubarray2(nums []int) int {
	stack := make([]int, 0)
	//单调连续递减栈
	left, right := 0, len(nums)-1
	stack = append(stack, len(nums)-1)
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[stack[len(stack)-1]] {
			for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
				stack = stack[0 : len(stack)-1]
			}
			if len(stack) == 0 {
				break
			}
		}
		if i == stack[len(stack)-1]-1 && nums[i] <= nums[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	right = right - len(stack)
	stack2 := append([]int{}, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[stack2[len(stack2)-1]] {
			for len(stack2) > 0 && nums[i] < nums[stack2[len(stack2)-1]] {
				stack2 = stack2[0 : len(stack2)-1]
			}
			if len(stack2) == 0 {
				break
			}
		}
		if i == stack2[len(stack2)-1]+1 {
			if nums[i] >= nums[stack2[len(stack2)-1]] {
				stack2 = append(stack2, i)
			}
		}
	}
	left += len(stack2)
	//fmt.Println(right,left,stack,stack2)
	if left >= right {
		return 0
	}
	return right - left + 1
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
