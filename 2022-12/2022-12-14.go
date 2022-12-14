package main

import "sort"

// 并查集模板
func findParent(a int, parent []int) int {
	if a == parent[a] {
		return a
	}
	parent[a] = findParent(parent[a], parent)
	return parent[a]
}
func merge(a, b int, parent []int) {
	pa, pb := findParent(a, parent), findParent(b, parent)
	parent[pa] = pb
}
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	res := make([]bool, len(queries))
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})
	qs := make([][]int, len(queries))
	for i := 0; i < len(queries); i++ {
		qs[i] = []int{i, queries[i][0], queries[i][1], queries[i][2]}
	}
	sort.Slice(qs, func(i, j int) bool {
		return qs[i][3] < qs[j][3]
	})
	j := 0
	for i := 0; i < len(queries); i++ {
		for j < len(edgeList) && edgeList[j][2] < qs[i][3] {
			merge(edgeList[j][0], edgeList[j][1], parent)
			j++
		}
		res[qs[i][0]] = findParent(qs[i][1], parent) == findParent(qs[i][2], parent)
	}
	return res
}
