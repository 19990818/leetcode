package main

import (
	"math"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	s := []int{0}
	m := make(map[int]int)
	res := new(ListNode)
	res.Next = head
	sum := 0
	for head != nil {
		sum += head.Val
		if sum == 0 {
			res.Next = head.Next
			s = []int{0}
			m = make(map[int]int)
			m[0] = 1
			head = head.Next
			continue
		}
		if _, ok := m[sum]; ok {
			i := 0
			for cur := res; cur != nil; cur = cur.Next {
				if s[i] == sum {
					cur.Next = head.Next
					for j := i; j < len(s); j++ {
						delete(m, s[j])
					}
					s = s[0:i]
					break
				}
				i++
			}

		}
		m[sum] = 1
		s = append(s, sum)
		head = head.Next
	}
	return res.Next
}

func findNonMinOrMax(nums []int) int {
	m := make(map[int]int)
	ma, mi := 0, 1000
	for _, num := range nums {
		m[num] = 1
		ma = max(ma, num)
		mi = min(mi, num)
	}
	for k := range m {
		if k != ma && k != mi {
			return k
		}
	}
	return -1
}

// 只要不遇到a全部需要进行替换
func smallestString(s string) string {
	var res strings.Builder
	i := 0
	for i < len(s) && s[i] == 'a' {
		i++
	}
	if i == len(s) {
		res.WriteString(s[0 : len(s)-1])
		res.WriteByte('z')
		return res.String()
	}
	res.WriteString(s[0:i])
	for ; i < len(s); i++ {
		if s[i] != 'a' {
			res.WriteByte(byte(s[i] - 1))
		} else {
			res.WriteString(s[i:])
			break
		}

	}
	return res.String()
}

func minCost(nums []int, x int) int64 {
	res := math.MaxInt64
	dp := make([][]int, len(nums))
	for i := range nums {
		dp[i] = getSigleMinCost(nums, x, i)
	}
	for cnt := 0; cnt < len(nums); cnt++ {
		temp := cnt * x
		for _, v := range dp {
			temp += v[cnt]
		}
		res = min(res, temp)
	}
	return int64(res)
}

func getSigleMinCost(nums []int, x int, idx int) []int {
	dp := make([]int, len(nums))
	dp[0] = nums[idx]
	for i := idx + 1; i < idx+len(nums); i++ {
		dp[i-idx] = min(dp[i-idx-1], nums[i%len(nums)])
	}
	return dp
}

func isFascinating(n int) bool {
	s := strconv.Itoa(n) + strconv.Itoa(2*n) + strconv.Itoa(3*n)
	if len(s) != 9 {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s {
		if v == rune('0') {
			return false
		}
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = 1
	}
	return true
}

func longestSemiRepetitiveSubstring(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if check(s[i : j+1]) {
				res = max(res, j+1-i)
			}
		}
	}
	return res
}
func check(s string) bool {
	cnt := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
		}
	}
	return cnt <= 1
}
