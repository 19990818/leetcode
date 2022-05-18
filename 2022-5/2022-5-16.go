package main

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var inTree func(root *TreeNode, p *TreeNode) bool
	inTree = func(root, p *TreeNode) bool {
		if root == p {
			return true
		}
		if root == nil {
			return false
		}
		return inTree(root.Left, p) || inTree(root.Right, p)
	}
	if inTree(root, p) == false {
		return nil
	}
	res := make([]*TreeNode, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root)
		dfs(root.Right)
	}
	dfs(root)
	//fmt.Println(res)
	for i := 0; i < len(res)-1; i++ {
		if res[i] == p {
			return res[i+1]
		}
	}
	return nil
}

func minimumDeleteSum(s1 string, s2 string) int {
	sum1 := 0
	sum2 := 0
	for _, val := range s1 {
		sum1 += int(val)
	}
	for _, val2 := range s2 {
		sum2 += int(val2)
	}
	m, n := len(s1), len(s2)
	if m == 0 || n == 0 {
		return sum1 + sum2
	}
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		c1 := s1[i-1]
		for j := 1; j < n+1; j++ {
			c2 := s2[j-1]
			if c1 == c2 {
				dp[i][j] = dp[i-1][j-1] + int(c1)
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return sum1 + sum2 - 2*dp[m][n]
}

func maxProfit(prices []int, fee int) int {
	stack := make([]int, 0)
	ans := 0
	for _, val := range prices {
		if len(stack) == 0 {
			stack = append(stack, val)
		} else if val >= stack[len(stack)-1] {
			stack = append(stack, val)
		} else if val+fee < stack[len(stack)-1] {
			ans += max(0, stack[len(stack)-1]-stack[0]-fee)
			stack = []int{val}
		} else if val < stack[0] {
			stack = []int{val}
		}
	}
	ans += max(0, stack[len(stack)-1]-stack[0]-fee)
	return ans
}
