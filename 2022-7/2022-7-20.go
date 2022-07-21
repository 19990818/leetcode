package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
	m := make(map[*TreeNode]int)
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if m[root] == 1 {
			return true
		}
		if root == nil {
			return true
		}
		if root.Left == nil && root.Right == nil && root.Val == 0 {
			m[root] = 1
			return true
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		if root.Val == 0 && l && r {
			m[root] = 1
			return true
		}
		return false
	}
	dfs(root)
	// fmt.Println(m)
	var travel func(root *TreeNode) *TreeNode
	travel = func(root *TreeNode) *TreeNode {
		if m[root] == 1 {
			return nil
		}
		if root.Left != nil {
			root.Left = travel(root.Left)
		}
		if root.Right != nil {
			root.Right = travel(root.Right)
		}

		return root
	}
	return travel(root)
}
