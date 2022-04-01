package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, c int) int {
	if a > c {
		return a
	}
	return c
}
func findFrequentTreeSum(root *TreeNode) []int {
	ansMap := make(map[int]int)
	var dfsTree func(root *TreeNode) int
	dfsTree = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sum := root.Val
		sum += dfsTree(root.Left)
		sum += dfsTree(root.Right)
		ansMap[sum]++
		return sum
	}
	dfsTree(root)
	ans := make([]int, 0)
	maxCount := 0
	for _, val := range ansMap {
		maxCount = max(maxCount, val)
	}
	for key, val := range ansMap {
		if val == maxCount {
			ans = append(ans, key)
		}
	}
	return ans
}

func findBottomLeftValue(root *TreeNode) int {
	var orderTravel func(root *TreeNode) []*TreeNode
	orderTravel = func(root *TreeNode) []*TreeNode {
		queue := make([]*TreeNode, 0)
		queue = append(queue, root)
		var ans []*TreeNode
		for {
			ans = append([]*TreeNode{}, queue...)
			nextQueue := make([]*TreeNode, 0)
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				if cur.Left != nil {
					nextQueue = append(nextQueue, cur.Left)
				}
				if cur.Right != nil {
					nextQueue = append(nextQueue, cur.Right)
				}
			}
			if len(nextQueue) == 0 {
				break
			}
			queue = nextQueue
		}
		return ans
	}
	lastChildren := orderTravel(root)
	return lastChildren[0].Val
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for {
		nextQueue := make([]*TreeNode, 0)
		maxElem := math.MinInt32
		for len(queue) > 0 {
			cur := queue[0]
			maxElem = max(maxElem, cur.Val)
			queue = queue[1:]
			if cur.Left != nil {
				nextQueue = append(nextQueue, cur.Left)
			}
			if cur.Right != nil {
				nextQueue = append(nextQueue, cur.Right)
			}
		}
		ans = append(ans, maxElem)
		if len(nextQueue) == 0 {
			break
		}
		queue = nextQueue
	}
	return ans
}
