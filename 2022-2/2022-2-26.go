package main

import (
	"math"
	"strings"
)

// 树的先序遍历实际上是首先是遍历根节点 我们将节点放入栈中 当遇到#表示栈顶节点
// 左边部分遍历完成 可以将该节点出栈开始右边部分的 最后肯定是以#结尾 不然不符合
// 如果栈中元素不存在 仍然有# 说明不符合规则 最后一个#不在考虑范围
func isValidSerialization(preorder string) bool {
	preorderArr := strings.Split(preorder, ",")
	if preorder == "#" {
		return true
	}
	if preorderArr[0] == "#" || len(preorderArr) < 3 || preorderArr[len(preorderArr)-1] != "#" {
		return false
	}
	stack := make([]string, 0)
	for idx := 0; idx < len(preorderArr)-1; idx++ {
		if len(stack) == 0 && preorderArr[idx] == "#" {
			return false
		} else if preorderArr[idx] == "#" {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, preorderArr[idx])
		}
	}
	return len(stack) == 0
}

// 此题没有周赛中那么复杂 只需要判断是否存在递增的三元组即可
// 周赛中是需要判断有多少个三元组 这是一种判断是否能实现和如何实现的区别
// 虽然看起来很像 但是本质上不是一个问题
func increasingTriplet(nums []int) bool {
	a, b := math.MaxInt64, math.MaxInt64
	for _, val := range nums {
		if val < a {
			a = val
		} else if val < b {
			b = val
		} else {
			return true
		}
	}
	return false
}

// 小偷抢劫问题 之前都是数组形式的 很容易得到动态规划表达式 面对这种问题 其本质
// 实际上还是递归问题 动态规划只是将递归的部分使用数组的形式进行保存，因为递归超时
// 的原因实际上就是对相同的问题会存在多次调用，因此将这些相同的问题严格控制在只让其
// 调用一次 就可以解决超时问题
func rob(root *TreeNode) int {
	dp := make(map[*TreeNode]int)
	if root == nil {
		return 0
	}
	return robRoot(root, dp)
}
func robRoot(root *TreeNode, dp map[*TreeNode]int) int {
	if root == nil {
		dp[root] = 0
		return 0
	}
	ans := root.Val
	if root.Left != nil {
		if _, ok := dp[root.Left.Left]; !ok {
			dp[root.Left.Left] = robRoot(root.Left.Left, dp)
		}
		ans += dp[root.Left.Left]
		if _, ok := dp[root.Left.Right]; !ok {
			dp[root.Left.Right] = robRoot(root.Left.Right, dp)
		}
		ans += dp[root.Left.Right]
	}
	if root.Right != nil {
		if _, ok := dp[root.Right.Left]; !ok {
			dp[root.Right.Left] = robRoot(root.Right.Left, dp)
		}
		ans += dp[root.Right.Left]
		if _, ok := dp[root.Right.Right]; !ok {
			dp[root.Right.Right] = robRoot(root.Right.Right, dp)
		}
		ans += dp[root.Right.Right]
	}
	ans = max(ans, robRoot(root.Left, dp)+robRoot(root.Right, dp))
	return ans
}
