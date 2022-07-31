package main

import (
	"math"
	"sort"
	"strings"
)

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 1
	queue := []*TreeNode{root}
	maxSum := math.MinInt64
	cnt := 1
	for {
		temp := []*TreeNode{}
		sum := 0
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			sum += cur.Val
			if cur.Left != nil {
				temp = append(temp, cur.Left)
			}
			if cur.Right != nil {
				temp = append(temp, cur.Right)
			}
		}
		if sum > maxSum {
			ans = cnt
		}
		if len(temp) == 0 {
			break
		}
		queue = temp
		cnt++
	}
	return ans
}

func permute(nums []int) [][]int {
	//不含重复数字
	n := len(nums)
	m := make(map[int]int)
	ans := make([][]int, 0)
	mid := make([]int, 0)
	var dfs func(l int)
	dfs = func(l int) {
		if l == n {
			temp := append([]int{}, mid...)
			ans = append(ans, temp)
		}
		for _, val := range nums {
			if m[val] == 1 {
				continue
			}
			m[val] = 1
			mid = append(mid, val)
			dfs(l + 1)
			mid = mid[:len(mid)-1]
			m[val] = 0
		}
	}
	dfs(0)
	return ans
}

func rotate(matrix [][]int) {
	//先按照对角线进行交换 再按照行进行交换
	var swap1, swap2 func(a [][]int)
	//对角线交换
	swap1 = func(a [][]int) {
		n := len(a)
		for i := 0; i < n; i++ {
			for j := 0; j < n-i; j++ {
				a[i][j], a[n-1-j][n-1-i] = a[n-1-j][n-1-i], a[i][j]
			}
		}
	}
	swap2 = func(a [][]int) {
		n := len(a)
		for up, down := 0, n-1; up < down; up, down = up+1, down-1 {
			for j := 0; j < n; j++ {
				a[up][j], a[down][j] = a[down][j], a[up][j]
			}
		}
	}
	swap1(matrix)
	swap2(matrix)
}

func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	m := make(map[string][]string)
	var sortString func(a string) string
	sortString = func(a string) string {
		b := make([]int, 26)
		for _, val := range a {
			b[int(val-'a')]++
		}
		var res strings.Builder
		for i := 0; i < 26; i++ {
			for cnt := 0; cnt < b[i]; cnt++ {
				res.WriteByte(byte(i + 'a'))
			}
		}
		return res.String()
	}
	for _, val := range strs {
		m[sortString(val)] = append(m[sortString(val)], val)
	}
	for _, val := range m {
		ans = append(ans, val)
	}
	return ans
}

// #53 最大数组和
func maxSubArray(nums []int) int {
	//实际上可以变为子问题 我们取前面的值 或者不取前面的值
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = max(0, dp[i-1]) + nums[i]
		ans = max(dp[i], ans)
	}
	return ans
}

// #55 跳跃游戏
func canJump(nums []int) bool {
	//记录当前最大值可以到达何处 更新最大值
	cur := 0
	for i := range nums {
		if i > cur {
			return false
		}
		cur = max(cur, i+nums[i])
	}
	return true
}

// #56 合并区间
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//fmt.Println(intervals)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		tail := res[len(res)-1]
		if intervals[i][0] <= tail[1] {
			res[len(res)-1][1] = max(tail[1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

// #62 不同路径
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// #64 最小路径和
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// #70 爬楼梯
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// #75 颜色分类
func sortColors(nums []int) {
	cnt := make([]int, 3)
	for _, num := range nums {
		cnt[num]++
	}
	sum := make([]int, 2)
	sum[0] = cnt[0]
	sum[1] = cnt[0] + cnt[1]
	for i := len(nums) - 1; i >= 0; i-- {
		if i >= sum[1] {
			nums[i] = 2
		} else if i >= sum[0] {
			nums[i] = 1
		} else {
			nums[i] = 0
		}
	}
}

// #78 子集
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	mid := make([]int, 0)
	var dfs func(cur int)
	ans = append(ans, []int{})
	dfs = func(cur int) {
		if cur == len(nums) {
			return
		}
		for i := cur; i < len(nums); i++ {
			mid = append(mid, nums[i])
			temp := append([]int{}, mid...)
			ans = append(ans, temp)
			dfs(i + 1)
			mid = mid[0 : len(mid)-1]
		}
	}
	dfs(0)
	return ans
}

// #79 单词搜索
func exist(board [][]byte, word string) bool {
	towards := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	m, n := len(board), len(board[0])
	travel := make(map[int]int)
	var dfs func(x, y int, w string) bool
	dfs = func(x, y int, w string) bool {
		if len(w) == 1 {
			return string(board[x][y]) == w
		}
		if board[x][y] == w[0] {
			for _, to := range towards {
				x2, y2 := x+to[0], y+to[1]
				if x2 < m && x2 >= 0 && y2 >= 0 && y2 < n && travel[x2*n+y2] == 0 {
					travel[x2*n+y2] = 1
					ans := dfs(x2, y2, w[1:])
					travel[x2*n+y2] = 0
					if ans {
						return true
					}
				}
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				travel[i*n+j] = 1
				ans := dfs(i, j, word)
				travel[i*n+j] = 0
				if ans {
					return true
				}
			}
		}
	}
	return false
}

// #94 中序遍历
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	//递归写法
	// ans = append(ans, inorderTraversal(root.Left)...)
	// ans = append(ans, root.Val)
	// ans = append(ans, inorderTraversal(root.Right)...)
	// return ans
	// 迭代写法
	cur := root
	stack := make([]*TreeNode, 0)
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			v := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			ans = append(ans, v.Val)
			cur = v.Right
		}
	}
	return ans
}

// #96 不同的二叉搜索树
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// #98 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//中序遍历二叉搜索树
	arr := inorderTraversal(root)
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

// #101 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var Symmetric func(r1, r2 *TreeNode) bool
	Symmetric = func(r1, r2 *TreeNode) bool {
		if r1 == nil || r2 == nil {
			return r1 == r2
		}
		if r1.Val != r2.Val {
			return false
		}
		return Symmetric(r1.Left, r2.Right) && Symmetric(r2.Left, r1.Right)
	}
	return Symmetric(root.Left, root.Right)
}

// #102 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for {
		temp := make([]int, 0)
		tQueue := make([]*TreeNode, 0)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			temp = append(temp, cur.Val)
			if cur.Left != nil {
				tQueue = append(tQueue, cur.Left)
			}
			if cur.Right != nil {
				tQueue = append(tQueue, cur.Right)
			}
		}
		ans = append(ans, temp)
		if len(tQueue) == 0 {
			break
		}
		queue = tQueue
	}
	return ans
}

// #104 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// #128 最长连续序列
func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val] = 1
	}
	ans := 0
	travel := make(map[int]int)
	for key := range m {
		temp := key
		if travel[temp] == 0 {
			cnt := 0
			for m[temp] != 0 {
				cnt++
				travel[temp] = 1
				temp--
			}
			ans = max(ans, cnt)
		}
	}
	return ans
}

// #139 单词拆分
func wordBreak(s string, wordDict []string) bool {
	// 表示以i结束的字符串是否满足条件
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, val := range wordDict {
			if len(val) <= i && s[i-len(val):i] == val {
				dp[i] = dp[i] || dp[i-len(val)]
			}
		}
	}
	return dp[len(s)]
}

// #105 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	idx := 0
	for i, val := range inorder {
		if val == preorder[0] {
			idx = i
			break
		}
	}
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}

// #114 二叉树展开为链表
func flatten(root *TreeNode) {
	var findNext func(t *TreeNode) *TreeNode
	findNext = func(t *TreeNode) *TreeNode {
		if t == nil {
			return nil
		}
		if t.Left == nil && t.Right == nil {
			return t
		}
		if t.Right != nil {
			return findNext(t.Right)
		}
		return findNext(t.Left)
	}
	for root != nil {
		if root.Left != nil {
			temp := root.Right
			tail := findNext(root.Left)
			tail.Right = temp
			root.Right, root.Left = root.Left, nil
		}
		root = root.Right
	}
}
