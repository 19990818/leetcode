package main

import (
	"strconv"
	"strings"
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ans := make([]*TreeNode, 0)
	m := make(map[string]int)
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		var temp strings.Builder
		temp.WriteString("{")
		temp.WriteString(strconv.Itoa(root.Val))
		temp.WriteString("}")
		temp.WriteString(dfs(root.Left))
		temp.WriteString(dfs(root.Right))
		m[temp.String()]++
		if m[temp.String()] == 2 {
			ans = append(ans, root)
		}
		return temp.String()
	}
	dfs(root)
	return ans
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	ans := new(TreeNode)
	maxNum, pos := nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			pos = i
		}
	}
	ans.Val = maxNum
	ans.Left = constructMaximumBinaryTree(nums[0:pos])
	ans.Right = constructMaximumBinaryTree(nums[pos+1:])
	return ans
}

func printTree(root *TreeNode) [][]string {
	var getHight func(root *TreeNode) int
	getHight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + max(getHight(root.Left), getHight(root.Right))
	}
	hight := getHight(root)
	ans := make([][]string, hight)
	for i := 0; i < hight; i++ {
		ans[i] = make([]string, pow(2, hight)-1)
		for j := 0; j < len(ans[i]); j++ {
			ans[i][j] = ""
		}
	}
	var dfs func(root *TreeNode, r, c int)
	dfs = func(root *TreeNode, r, c int) {
		if root == nil {
			return
		}
		ans[r][c] = strconv.Itoa(root.Val)
		dfs(root.Left, r+1, c-pow(2, hight-r-2))
		dfs(root.Right, r+1, c+pow(2, hight-r-2))
	}
	dfs(root, 0, (pow(2, hight)-1)/2)
	return ans
}
func pow(a int, b int) int {
	ans := 1
	for b > 0 {
		if b%2 == 1 {
			ans *= a
		}
		b = b / 2
		a *= a
	}
	return ans
}
