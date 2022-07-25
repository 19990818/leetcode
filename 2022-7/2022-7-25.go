package main

import (
	"sort"
	"strings"
)

func repeatedCharacter(s string) byte {
	cnt := make(map[rune]int)
	for _, val := range s {
		cnt[val]++
		if cnt[val] == 2 {
			return byte(val)
		}
	}
	return ' '
}

func equalPairs(grid [][]int) int {
	n := len(grid)
	ans := 0
	m1, m2 := make(map[string]int), make(map[string]int)
	for i := 0; i < n; i++ {
		var temp1, temp2 strings.Builder
		for j := 0; j < n; j++ {
			temp1.WriteByte(byte('0' + grid[i][j]))
			temp2.WriteByte(byte('0' + grid[j][i]))
		}
		m1[temp1.String()]++
		m2[temp2.String()]++
	}
	//fmt.Println(m1,m2)
	for key := range m1 {
		if m2[key] != 0 {
			ans += m1[key] * m2[key]
		}
	}
	return ans
}

func countExcellentPairs(nums []int, k int) int64 {
	m := make(map[int]int)
	var cnt1 func(a int) int
	cnt := make([]int, 0)
	cnt1 = func(a int) int {
		res := 0
		for a > 0 {
			res += a % 2
			a /= 2
		}
		return res
	}
	for _, val := range nums {
		if m[val] == 0 {
			cnt = append(cnt, cnt1(val))
			m[val] = 1
		}

	}
	sort.Ints(cnt)
	ans := int64(0)
	left, right := 0, len(cnt)-1
	//fmt.Println(cnt)
	for left <= right {
		for left <= right && cnt[left]+cnt[right] < k {
			left++
		}
		//fmt.Println(left,right)
		ans += int64(max((right-left+1)*2-1, 0))
		right--
	}
	return ans
}
