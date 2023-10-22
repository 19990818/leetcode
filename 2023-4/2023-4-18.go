package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	// 得到子树的最大值和最小值
	res := 0
	var dfs func(r *TreeNode) (int, int)
	dfs = func(r *TreeNode) (int, int) {
		if r == nil {
			return 0, 0
		}
		ma1, mi1 := dfs(r.Left)
		ma2, mi2 := dfs(r.Right)
		ma := max(ma1, ma2)
		mi := min(mi1, mi2)
		res = max(res, max(abs(ma-r.Val), abs(r.Val-mi)))
		return ma, mi
	}
	dfs(root)
	return res
}
