package main

import (
	"sort"
	"strconv"
)

func rearrangeBarcodes(barcodes []int) []int {
	m := make(map[int]int)
	for _, v := range barcodes {
		m[v]++
	}
	ma := 0
	mv := barcodes[0]
	for k, v := range m {
		if v > ma {
			mv = k
			ma = v
		}
	}
	temp := make([][]int, ma)
	for i := range temp {
		temp[i] = []int{mv}
	}
	i := 0
	for k, v := range m {
		if v > 0 {
			temp[i] = append(temp[i], k)
			v--
			i = (i + 1) % ma
		}
	}
	res := make([]int, 0)
	for _, t := range temp {
		res = append(res, t...)
	}
	return res
}

func countSeniors(details []string) int {
	res := 0
	for _, v := range details {
		age, _ := strconv.Atoi(v[11:13])
		if age > 60 {
			res++
		}
	}
	return res
}

func matrixSum(nums [][]int) int {
	res := 0
	for i := range nums {
		sort.Ints(nums[i])
	}
	for j := 0; j < len(nums[0]); j++ {
		t := 0
		for i := 0; i < len(nums); i++ {
			t = max(t, nums[i][j])
		}
		res += t
	}
	return res
}

func maximumOr(nums []int, k int) int64 {
	// 实际上只会对一个数字进行加倍，因为其最高位移动到一个新的上界
	pre, suf := make([]int, len(nums)+1), make([]int, len(nums)+1)
	pre[0] = 0
	suf[len(nums)] = 0
	for i := 0; i < len(nums); i++ {
		pre[i+1] = pre[i] | nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		suf[i] = suf[i+1] | nums[i]
	}
	res := 0
	for i, v := range nums {
		res = max(res, pre[i]|suf[i+1]|(v*(1<<k)))
	}
	return int64(res)
}

func sumOfPower(nums []int) int {
	sort.Ints(nums)
	mod := int(1e9 + 7)
	temp := nums[0]
	res := ((nums[0] * nums[0]) % mod * nums[0]) % mod
	for i := 1; i < len(nums); i++ {
		m := (nums[i] * nums[i]) % mod
		res = (res + (m*(temp+nums[i]))%mod) % mod
		temp = (temp*2 + nums[i]) % mod
	}
	return res
}
