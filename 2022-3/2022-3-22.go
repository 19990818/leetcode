package main

import "sort"

type interval [][]int

func (m interval) Less(i, j int) bool {
	return m[i][1] < m[j][1]
}
func (m interval) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m interval) Len() int {
	return len(m)
}
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Sort(interval(intervals))
	ans := 0
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= cur[1] {
			cur = intervals[i]
		} else {
			ans++
		}
	}
	return ans
}

func findRightInterval(intervals [][]int) []int {
	mStart := make(map[int]int)
	start := make([]int, 0)
	for idx, val := range intervals {
		mStart[val[0]] = idx
		start = append(start, val[0])
	}
	sort.Ints(start)
	ans := make([]int, 0)
	for _, val := range intervals {
		temp := getMinBigger(val[1], start)
		if temp == -1 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, mStart[start[temp]])
		}
	}
	return ans
}
func getMinBigger(target int, src []int) int {
	if target > src[len(src)-1] {
		return -1
	}
	left, right := 0, len(src)-1
	for left < right {
		mid := (right-left)>>1 + left
		if src[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := 0
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	ans += getFixed(root, targetSum)
	return ans
}
func getFixed(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := 0
	if root.Val == targetSum {
		ans++
	}
	ans += getFixed(root.Left, targetSum-root.Val)
	ans += getFixed(root.Right, targetSum-root.Val)
	return ans
}
