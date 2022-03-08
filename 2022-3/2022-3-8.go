package main

import "sort"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	left, right := nums1[0]+nums2[0], nums1[len(nums1)-1]+nums2[len(nums2)-1]
	m, n := len(nums1), len(nums2)
	pairSum := left + sort.Search(right-left, func(sum int) bool {
		sum += left
		i, j := 0, n-1
		count := 0
		for i < m && j >= 0 {
			if nums1[i]+nums2[j] <= sum {
				i++
				count += j + 1
			} else {
				j--
			}
		}
		return count >= k
	})
	//fmt.Println(pairSum)
	j := n - 1
	ans := make([][]int, 0)
	for _, num1 := range nums1 {
		for j >= 0 && num1+nums2[j] >= pairSum {
			j--
		}
		for _, num2 := range nums2[:j+1] {
			ans = append(ans, []int{num1, num2})
			if len(ans) == k {
				return ans
			}
		}
	}
	//处理得到等于k小的值
	j = n - 1
	temp := n - 1
	for _, num1 := range nums1 {
		j = temp
		for j >= 0 && num1+nums2[j] > pairSum {
			j--
		}
		temp = j
		for j >= 0 {
			if num1+nums2[j] == pairSum {
				ans = append(ans, []int{num1, nums2[j]})
				j--
				if len(ans) == k {
					return ans
				}
			} else {
				break
			}
		}
	}
	return ans
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := (right-left)>>1 + left
		if checkNum(matrix, mid, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func checkNum(matrix [][]int, mid, k int) bool {
	i, j := 0, len(matrix[0])-1
	count := 0
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] <= mid {
			i++
			count += j + 1
		} else {
			j--
		}
	}
	return count >= k
}

// 为什么n要从后面开始计算 因为98行中k+1>i 若从小到大遍历无法获得正确的dp[k+1][j]的值
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j < n+1; j++ {
			dp[i][j] = dp[i][j-1] + j
			temp := 0
			for k := i; k < j; k++ {
				temp = max(dp[i][k-1], dp[k+1][j]) + k
				if temp < dp[i][j] {
					dp[i][j] = temp
				}
			}
		}
	}
	return dp[1][n]
}
