package main

func maxSumBST(root *TreeNode) int {
	var dfs func(root *TreeNode) btree
	res := 0
	dfs = func(root *TreeNode) btree {
		if root == nil {
			return btree{-1e5, 1e5, 0, true}
		}
		btl, btr := dfs(root.Left), dfs(root.Right)
		if btl.isBST && btr.isBST && btl.maxv < root.Val && root.Val < btr.minv {
			temp := btl.sum + btr.sum + root.Val
			res = max(res, temp)
			//fmt.Println(temp)
			return btree{max(btr.maxv, root.Val), min(btl.minv, root.Val), temp, true}
		}
		return btree{0, 0, 0, false}
	}
	dfs(root)
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type btree struct {
	maxv  int
	minv  int
	sum   int
	isBST bool
}

func circularGameLosers(n int, k int) []int {
	m := make(map[int]int)
	s := 0
	i := 1
	for m[s+1] == 0 {
		m[s+1] = 1
		s = (s + i*k) % n
		i++
	}
	res := make([]int, 0)
	for i := 1; i <= n; i++ {
		if m[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}

func doesValidArrayExist(derived []int) bool {
	s := 0
	for i := 0; i < len(derived)-1; i++ {
		s ^= derived[i]
	}
	return (s ^ 0) == derived[len(derived)-1]
}
