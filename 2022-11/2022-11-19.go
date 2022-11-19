package main

import (
	"sort"
)

func sumSubseqWidths(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	pow := make([]int, n)
	mod := int(1e9 + 7)
	pow[0] = 1
	for i := 1; i < n; i++ {
		pow[i] = (pow[i-1] * 2) % mod
	}
	res := 0
	for i := 0; i < n; i++ {
		temp := (pow[i] - pow[n-i-1]) % mod
		res = (res + temp) % mod
	}
	return res
}

func minimumOperations(root *TreeNode) int {
	q := []*TreeNode{root}
	ans := 0
	for len(q) > 0 {
		n := len(q)
		temp := q
		q = nil
		a := make([]int, n)
		for i, v := range temp {
			a[i] = v.Val
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
		id := make([]int, n)
		for i := range id {
			id[i] = i
		}
		ans += n
		sort.Slice(id, func(i, j int) bool { return a[id[i]] < a[id[j]] })
		vis := make([]bool, n)
		for i := 0; i < n; i++ {
			if !vis[i] {
				for p := i; !vis[p]; p = id[p] {
					vis[p] = true
				}
				ans--
			}
		}
	}
	return ans
}
