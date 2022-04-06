package main

import (
	"math"
	"math/rand"
)

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum < target {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	a := (sum + target) / 2
	count := 0
	var dfs func(nums []int, a int)
	dfs = func(nums []int, a int) {
		if a == 0 {
			count++
		}
		for i := 0; i < len(nums); i++ {
			dfs(nums[i+1:], a-nums[i])
		}
	}
	dfs(nums, a)
	return count
}

type SolutionRec struct {
	recs [][]int
	maxX int
	maxY int
	minX int
	minY int
}

func ConstructorRec(rects [][]int) SolutionRec {
	maxX, maxY := math.MinInt32, math.MinInt32
	minX, minY := math.MaxInt32, math.MaxInt32
	for _, val := range rects {
		maxX = max(maxX, val[2])
		maxY = max(maxY, val[3])
		minX = min(minX, val[0])
		minY = min(minY, val[1])
	}
	return SolutionRec{rects, maxX, maxY, minX, minY}
}

func (this *SolutionRec) Pick() []int {
	var isInPool func(x, y int, rects [][]int) bool
	isInPool = func(x, y int, rects [][]int) bool {
		for _, val := range rects {
			if x >= val[0] && x <= val[2] && y >= val[1] && y <= val[3] {
				return true
			}
		}
		return false
	}
	var x, y int
	for {
		x, y = rand.Intn(this.maxX-this.minX)+this.minX, rand.Intn(this.maxY-this.minY)+this.minY
		if isInPool(x, y, this.recs) {
			break
		}
	}
	return []int{x, y}
}

func findDiagonalOrder(mat [][]int) []int {
	count := 0
	ans := make([]int, 0)
	maxCount := len(mat) + len(mat[0]) - 1
	for count < maxCount {
		if count%2 == 0 {
			for i := count; i >= 0; i-- {
				if i < len(mat) && count-i < len(mat[0]) {
					ans = append(ans, mat[i][count-i])
				}
			}
		} else {
			for j := count; j >= 0; j-- {
				if count-j < len(mat) && j < len(mat[0]) {
					ans = append(ans, mat[count-j][j])
				}

			}
		}
		count++
	}
	return ans
}

func nextGreaterElements(nums []int) []int {
	nums = append(nums, nums...)
	ans := make([]int, len(nums))
	s := make([]int, 0)
	s = append(s, 0)
	for i := 1; i < len(nums); i++ {
		for len(s) > 0 && nums[i] > nums[s[len(s)-1]] {
			cur := s[len(s)-1]
			s = s[0 : len(s)-1]
			ans[cur] = nums[i]
		}
		s = append(s, i)
	}
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[0 : len(s)-1]
		ans[cur] = -1
	}
	return ans[0 : len(nums)/2]
}
