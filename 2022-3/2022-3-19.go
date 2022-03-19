package main

import "sort"

type Node4Ary struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node4Ary
	TopRight    *Node4Ary
	BottomLeft  *Node4Ary
	BottomRight *Node4Ary
}

func construct(grid [][]int) *Node4Ary {
	if isFullSame(grid) {
		return &Node4Ary{grid[0][0] == 1, true, nil, nil, nil, nil}
	}
	m := len(grid)
	ans := &Node4Ary{Val: true, IsLeaf: false}
	arrTopLeft := splitArr(grid, 0, 0)
	arrTopRight := splitArr(grid, 0, m/2)
	arrBottomLeft := splitArr(grid, m/2, 0)
	arrBottomRight := splitArr(grid, m/2, m/2)
	ans.TopLeft = construct(arrTopLeft)
	ans.TopRight = construct(arrTopRight)
	ans.BottomLeft = construct(arrBottomLeft)
	ans.BottomRight = construct(arrBottomRight)
	return ans
}
func splitArr(grid [][]int, row, col int) [][]int {
	length := len(grid)
	ans := make([][]int, length/2)
	for i := 0; i < length/2; i++ {
		ans[i] = make([]int, length/2)
	}
	for i := row; i < row+length/2; i++ {
		for j := col; j < col+length/2; j++ {
			ans[i][j] = grid[i][j]
		}
	}
	return ans
}
func isFullSame(grid [][]int) bool {
	flag := grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != flag {
				return false
			}
		}
	}
	return true
}

type NodeNAry struct {
	Val      int
	Children []*NodeNAry
}

func levelOrder(root *NodeNAry) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([][]*NodeNAry, 0)
	queue = append(queue, []*NodeNAry{root})
	i := 0
	for len(queue) > i {
		temp := make([]*NodeNAry, 0)
		subAns := make([]int, 0)
		for len(queue[i]) > 0 {
			cur := queue[i][0]
			queue[i] = queue[i][1:]
			subAns = append(subAns, cur.Val)
			for _, val := range cur.Children {
				temp = append(temp, val)
			}
		}
		if len(temp) != 0 {
			queue = append(queue, temp)
		}
		i++
		ans = append(ans, subAns)
	}
	return ans
}

type NodeMultilevelList struct {
	Val   int
	Prev  *NodeMultilevelList
	Next  *NodeMultilevelList
	Child *NodeMultilevelList
}

func flatten(root *NodeMultilevelList) *NodeMultilevelList {
	cur := root
	for cur != nil {
		curNex := cur.Next
		if cur.Child != nil {
			temp := flatten(cur.Child)
			cur.Next = temp
			temp.Prev = cur
			cur.Child = nil
			tail := temp
			for ; tail.Next != nil; tail = tail.Next {
			}
			tail.Next = curNex
			if curNex != nil {
				curNex.Prev = tail
			}
		}
		cur = curNex
	}
	return root
}

func divideArray(nums []int) bool {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	for _, val := range m {
		if val%2 == 1 {
			return false
		}
	}
	return true
}
func maximumSubsequenceCount(text string, pattern string) int64 {
	last := 0
	for _, val := range text {
		if byte(val) == pattern[1] {
			last++
		}
	}
	pre := 0
	sum := int64(0)
	for idx := range text {
		if text[idx] == pattern[0] {
			pre++
		} else if text[idx] == pattern[1] {
			sum += int64(pre)
		}
	}
	if pattern[0] == pattern[1] {
		ans := int64(0)
		ans = int64(pre-1)*int64(pre)/2 + int64(pre)
		return ans
	}
	if last < pre {
		return sum + int64(pre)
	}
	return sum + int64(last)
}

func halveArray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	sum := 0
	mid := make([]float64, 0)
	for _, val := range nums {
		sum += val
		mid = append(mid, float64(val))
	}
	var target float64
	target = float64(sum) / 2
	sort.Float64s(mid)
	op := 0
	sumMinus := float64(0)
	sort.Float64s(mid)
	left := 1
	right := 1
	for sumMinus < target {
		if mid[len(mid)-left] >= mid[len(mid)-right] {
			sumMinus += mid[len(mid)-left] / 2
			mid[len(mid)-left] /= 2
			if left < len(mid)-1 && mid[len(mid)-1-left] > mid[len(mid)-left] {
				left++
			}
			if right == len(mid) {
				right = 1
			}

		} else {
			sumMinus += mid[len(mid)-right] / 2
			mid[len(mid)-right] /= 2
			if right < len(mid) && mid[len(mid)-right-1] > mid[len(mid)-right] {
				right++
			}
			if left == len(mid) {
				left = 1
			}
		}
		//fmt.Println(mid,sumMinus)
		op++

	}
	return op
}
