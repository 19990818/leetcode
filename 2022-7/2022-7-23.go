package main

import "math"

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	// 找到每一个的连接是否存在
	if len(nums) == 1 {
		return true
	}
	cnt := len(nums) - 1
	type link struct {
		cur  int
		next int
	}
	m := make(map[link]int)
	for i := 1; i < len(nums); i++ {
		m[link{nums[i-1], nums[i]}] = 1
	}
	//注意给到的seq总和数量限制在1e5因此可以进行这种双重循环
	for _, seq := range sequences {
		for i := 1; i < len(seq); i++ {
			if m[link{seq[i-1], seq[i]}] == 1 {
				cnt--
				if cnt == 0 {
					return true
				}
				m[link{seq[i-1], seq[i]}] = 0
			}
		}
	}
	return false
}

func bestHand(ranks []int, suits []byte) string {
	m := make(map[byte]int)
	m2 := make(map[int]int)
	for _, val := range suits {
		m[val] = 1
	}
	if len(m) == 1 {
		return "Flush"
	}
	for _, val := range ranks {
		m2[val]++
	}
	ans := 0
	for _, val := range m2 {
		if val > ans {
			ans = val
		}
	}
	if ans == 1 {
		return "High Card"
	}
	if ans == 2 {
		return "Pair"
	}
	return "Three of a Kind"
}

func zeroFilledSubarray(nums []int) int64 {
	ans := int64(0)
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt++
		} else {
			ans += int64((cnt + 1) * cnt / 2)
			cnt = 0
		}
	}
	if cnt > 0 {
		ans += int64((cnt + 1) * cnt / 2)
	}
	return ans
}

func shortestSequence(rolls []int, k int) int {
	last := make(map[int]int)
	lastPos := make(map[int]int)
	for i := len(rolls) - 1; i >= 0; i-- {
		if _, ok := last[rolls[i]]; !ok {
			last[rolls[i]] = i
			lastPos[i] = 1
		}
	}
	if len(last) < k {
		return 1
	}
	//找到每个
	ans := math.MaxInt64
	cnt := make(map[int]int)
	x := 0
	for i, val := range rolls {
		if lastPos[i] == 1 {
			if ans > x {
				ans = x
			}
		}
		cnt[val] = 1
		if len(cnt) == k {
			x++
			cnt = make(map[int]int)
		}
	}
	return ans + 2
}
