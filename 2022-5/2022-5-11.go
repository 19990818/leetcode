package main

import (
	"bytes"
	"strings"
)

func findRedundantConnection(edges [][]int) []int {
	//并查集
	n := len(edges)
	father := make([]int, n+1)
	for i := 1; i <= n; i++ {
		father[i] = i
	}
	var find func(u int) int
	find = func(u int) int {
		if u == father[u] {
			return u
		}
		father[u] = find(father[u])
		return father[u]
	}
	var insert func(u, v int)
	insert = func(u, v int) {
		u = find(u)
		v = find(v)
		if u == v {
			return
		}
		father[u] = v
	}
	for _, val := range edges {
		//fmt.Println(val[0],val[1])
		if find(val[0]) == find(val[1]) {
			return val
		} else {
			insert(val[0], val[1])
		}
	}
	return []int{}
}

func repeatedStringMatch(a string, b string) int {
	m := len(b) / len(a)
	aArr := bytes.Repeat([]byte(a), m)
	if strings.Contains(string(aArr), b) {
		return m
	}
	aArr = append(aArr, []byte(a)...)
	if strings.Contains(string(aArr), b) {
		return m + 1
	}
	aArr = append(aArr, []byte(a)...)
	if strings.Contains(string(aArr), b) {
		return m + 2
	}
	return -1
}

func longestUnivaluePath(root *TreeNode) int {
	var getRootMax func(root *TreeNode, val int) int
	getRootMax = func(root *TreeNode, val int) int {
		if root == nil {
			return 0
		}
		if root.Val != val {
			return 0
		}
		return max(1+getRootMax(root.Left, val), 1+getRootMax(root.Right, val))
	}
	if root == nil {
		return 0
	}
	ans := getRootMax(root.Left, root.Val) + getRootMax(root.Right, root.Val)
	ans = max(ans, longestUnivaluePath(root.Left))
	ans = max(ans, longestUnivaluePath(root.Right))
	return ans
}
