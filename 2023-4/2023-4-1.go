package main

import (
	"sort"
	"strings"
)

func maskPII(s string) string {
	email := strings.Contains(s, "@")
	if email {
		return dealEmail(s)
	}
	return dealPhone(s)
}
func dealEmail(s string) string {
	arr := strings.Split(s, "@")
	sufix := strings.ToLower(arr[1])
	arr[0] = strings.ToLower(arr[0])
	prefix := string(arr[0][0]) + "*****" + string(arr[0][len(arr[0])-1])
	return prefix + "@" + sufix
}
func dealPhone(s string) string {
	var numStr strings.Builder
	for _, v := range s {
		if v <= '9' && v >= '0' {
			numStr.WriteRune(v)
		}
	}
	var res strings.Builder
	numS := numStr.String()
	if len(numS) > 10 {
		res.WriteByte('+')
		for i := 0; i < len(numS)-10; i++ {
			res.WriteByte('*')
		}
		res.WriteString("-")
	}
	res.WriteString("***-***-")
	res.WriteString(numS[len(numS)-4:])
	return res.String()
}

func minOperations(nums []int, queries []int) []int64 {
	sort.Ints(nums)
	sum := make([]int, len(nums)+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}
	n := len(nums)
	res := make([]int64, len(queries))
	for i, v := range queries {
		idx := sort.SearchInts(nums, v+1)
		pre := idx*v - sum[idx]
		suf := sum[n] - sum[idx] - (n-idx)*v
		res[i] = int64(pre + suf)
	}
	return res
}

func evenOddBit(n int) []int {
	odd, even := 0, 0
	for i := 32; i >= 0; i-- {
		if n >= 1<<i {
			n -= 1 << i
			if i%2 == 0 {
				odd++
			} else {
				even++
			}
		}
	}
	return []int{odd, even}
}
