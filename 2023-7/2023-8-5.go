package main

import "sort"

type edge struct {
	x   int
	y   int
	dis int
}

func minCostConnectPoints(points [][]int) int {

	n := len(points)
	edges := make([]edge, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, edge{i, j, abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])})
		}
	}
	parent := make([]int, n)
	dis := make([]int, n)
	for i := range points {
		parent[i] = i
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dis < edges[j].dis
	})
	cnt := 0
	for _, e := range edges {
		temp := merge(e, parent, dis)
		if temp != 0 {
			cnt++
		}
		if cnt == n-1 {
			return temp
		}
	}
	return 0
}
func find(a int, parent []int) int {
	for a != parent[a] {
		a = parent[a]
	}
	return parent[a]
}
func merge(e edge, parent, dis []int) int {
	rootx := find(e.x, parent)
	rooty := find(e.y, parent)
	if rootx == rooty {
		return 0
	}
	parent[rooty] = rootx
	dis[rootx] += dis[rooty] + e.dis
	return dis[rootx]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
