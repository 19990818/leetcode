package main

import "sort"

func unequalTriplets(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					ans++
				}
			}
		}
	}
	return ans
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	tArr := make([]int, 0)
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		tArr = append(tArr, r.Val)
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	sort.Ints(tArr)
	//fmt.Println(tArr)
	ans := make([][]int, len(queries))
	for i := range queries {
		p := sort.SearchInts(tArr, queries[i])
		maxNum, minNum := -1, -1
		if p < len(tArr) && tArr[p] == queries[i] {
			maxNum = tArr[p]
		} else if p > 0 {
			maxNum = tArr[p-1]
		}
		if p < len(tArr) {
			minNum = tArr[p]
		}
		ans[i] = []int{maxNum, minNum}
	}
	return ans
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads) + 1
	out := make([][]int, n)
	for _, v := range roads {
		out[v[0]] = append(out[v[0]], v[1])
		out[v[1]] = append(out[v[1]], v[0])
	}
	order := make([][]int, 0)
	vis := make([]bool, n)
	q := []int{0}
	vis[0] = true
	for len(q) > 0 {
		temp := q
		q = nil
		sub := make([]int, 0)
		for _, v := range temp {
			sub = append(sub, v)
			for _, v2 := range out[v] {
				if !vis[v2] {
					q = append(q, v2)
					vis[v2] = true
				}
			}
		}
		order = append(order, temp)
	}
    //fmt.Println(order)
	ans := int64(0)
	sum := make([]int, n)
	vis2 := make([]bool, n)
	for i := len(order) - 1; i > 0; i-- {
		for _, v := range order[i] {
			ans += int64(sum[v]+seats)/int64(seats)
			vis2[v] = true
			for _, v2 := range out[v] {
				if !vis2[v2] {
					sum[v2] += sum[v]+1
				}
			}
		}
	}
    //fmt.Println(sum)
	return ans
}