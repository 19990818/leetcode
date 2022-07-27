package main

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0
	cur1, cur2 := l1, l2
	var tail *ListNode
	for cur1 != nil && cur2 != nil {

		cur2.Val, c = (cur1.Val+cur2.Val+c)%10, (cur1.Val+cur2.Val+c)/10

		cur1 = cur1.Next
		if cur2.Next == nil {
			tail = cur2
		}
		cur2 = cur2.Next
	}
	for cur1 != nil {
		temp := new(ListNode)
		temp.Val, c = (cur1.Val+c)%10, (cur1.Val+c)/10
		tail.Next = temp
		tail = tail.Next
		cur1 = cur1.Next
	}
	for cur2 != nil {
		cur2.Val, c = (cur2.Val+c)%10, (cur2.Val+c)/10
		if cur2.Next == nil {
			tail = cur2
		}
		cur2 = cur2.Next
	}
	if c > 0 {
		temp := new(ListNode)
		temp.Val = c
		tail.Next = temp
	}
	return l2
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	left := -1
	ans := 0
	for idx, val := range s {
		if _, ok := m[val]; !ok {
			m[val] = idx
			ans = max(ans, idx-left)
		} else {
			old := m[val]
			for i := left + 1; i <= old; i++ {
				delete(m, rune(s[i]))
			}
			//fmt.Println(idx,old,m)
			ans = max(ans, idx-old)
			left = old
			m[val] = idx
		}
	}
	return ans
}

func longestPalindrome(s string) string {
	//dp[i][j]表示以i开始 以j结束的字符串是否为回文
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := 0; j <= i; j++ {
			dp[i][j] = true
		}
	}
	start, end := 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = dp[i+1][j-1] && (s[i] == s[j])
			if dp[i][j] && j-i > end-start {
				start, end = i, j
			}
		}
	}
	return s[start : end+1]
}

func maxArea(height []int) int {
	//使用贪心的策略 制约我们得到最大面积的较短的那条边
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		if height[left] >= height[right] {
			//fmt.Println(height[right],right-left)
			ans = max(ans, height[right]*(right-left))
			right--
		} else {
			ans = max(ans, height[left]*(right-left))
			left++
		}
	}
	return ans
}

func threeSum(nums []int) [][]int {
	//三个数定下两个
	if len(nums) < 3 {
		return [][]int{}
	}
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	ans := make([][]int, 0)
	if m[0] >= 3 {
		ans = append(ans, []int{0, 0, 0})
	}
	temp := make([]int, 0)
	for key := range m {
		temp = append(temp, key)
	}
	sort.Ints(temp)
	for i := 0; i < len(temp) && temp[i] < 0; i++ {
		for j := i; j < len(temp); j++ {
			if -(temp[i]+temp[j]) < temp[j] || temp[i] == temp[j] && temp[i] == 0 {
				continue
			}
			m[temp[i]]--
			m[temp[j]]--
			if m[-(temp[i]+temp[j])] > 0 && m[temp[i]] >= 0 && m[temp[j]] >= 0 {
				ans = append(ans, []int{temp[i], temp[j], -(temp[i] + temp[j])})
			}
			m[temp[i]]++
			m[temp[j]]++
		}
	}
	return ans
}
