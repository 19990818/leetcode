package main

import "strings"

func largestGoodInteger(num string) string {
	for i := 9; i >= 0; i-- {
		var temp strings.Builder
		for j := 0; j < 3; j++ {
			temp.WriteRune(rune(i + '0'))
		}
		if strings.Contains(num, temp.String()) {
			return temp.String()
		}
	}
	return ""
}

func averageOfSubtree(root *TreeNode) int {
	var sum func(root *TreeNode) (int, int)
	sum = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		leftSum, leftNum := sum(root.Left)
		rightSum, rightNum := sum(root.Right)
		return (root.Val + leftSum + rightSum), (1 + leftNum + rightNum)
	}
	ans := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		rootSum, rootNum := sum(root)
		if root.Val == rootSum/rootNum {
			ans++
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return ans
}

func countTexts(pressedKeys string) int {
	mod := int(1e9 + 7)
	cur := pressedKeys[0]
	ans := 1
	count := 1
	var get func(count int, maxNum int) int
	get = func(count, maxNum int) int {
		dp := make([]int, count+1)
		dp[0] = 1
		for i := 1; i <= count; i++ {
			for j := 1; j <= maxNum && i >= j; j++ {
				dp[i] = (dp[i-j] + dp[i]) % mod
			}
		}
		//fmt.Println(dp[count])
		return dp[count]
	}
	var maxNum int
	maxNum = 3
	if cur == '7' || cur == '9' {
		maxNum = 4
	}
	for i := 1; i < len(pressedKeys); i++ {
		if pressedKeys[i] != cur {
			ans = ans * get(count, maxNum) % mod
			cur = pressedKeys[i]
			maxNum = 3
			if cur == '7' || cur == '9' {
				maxNum = 4
			}
			count = 1
		} else {
			count++
		}
	}
	//fmt.Println(maxNum)
	ans = (ans * get(count, maxNum)) % mod
	return ans
}
