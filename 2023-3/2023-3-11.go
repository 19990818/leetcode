package main

import "sort"

func findLongestSubarray(array []string) []string {
	// 记录下当前前缀和的cnt数量，之后遇到的cnt数量相同为满足条件
	cur := 0
	m := make(map[int]int)
	start, end := 0, 0
	maxLen := 0
	m[0] = -1
	for i, v := range array {
		if v[0] <= '9' && v[0] >= '0' {
			cur += 1
		} else {
			cur -= 1
		}
		if _, ok := m[cur]; ok {
			if i-m[cur] > maxLen {
				start, end = m[cur], i
				maxLen = i - m[cur]
			}
		} else {
			m[cur] = i
		}
	}
	return array[start+1 : end+1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	levelArr := make([]int, 0)
	q := []*TreeNode{root}
	for len(q) > 0 {
		temp := q
		q = nil
		sum := 0
		for _, v := range temp {
			sum += v.Val
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
		levelArr = append(levelArr, sum)
	}
	sort.Ints(levelArr)
	return int64(levelArr[len(levelArr)-k])
}

func findValidSplit(nums []int) int {
	left, right := 0, 0
	for ; left <= right && right < len(nums); left++ {
		for j := len(nums) - 1; j >= right+1; j-- {
			if check(nums[left], nums[j]) {
				right = j
				break
			}
		}
	}
	if right >= len(nums)-1 {
		return -1
	}
	return right
}

func check(a, b int) bool {
	for b > 0 {
		a, b = b, a%b
	}
	return a > 1
}
