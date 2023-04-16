package main

import (
	"sort"
	"strconv"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	sum := make([]int, len(nums)+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}
	res := make([]int, len(queries))
	for i, v := range queries {
		res[i] = sort.SearchInts(sum, v)
	}
	return res
}

func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	s = s + s
	vis := make([]int, n)
	res := s
	for i := 0; vis[i] == 0; i = (i + b) % n {
		vis[i] = 1

		for j := 0; j < 10; j++ {
			for m := 0; m < 10; m++ {
				t := append([]byte{}, s[i:i+n]...)
				for p := 1; p < n; p += 2 {
					t[p] = byte('0' + (int(t[p]-'0')+a*j)%10)
				}
				for p := 0; p < n; p += 2 {
					t[p] = byte('0' + (int(t[p]-'0')+a*m)%10)
				}
				if string(t) < res {
					res = string(t)
				}
			}
			//fmt.Println(string(t),res)

		}
	}
	return res
}

func numDupDigitsAtMostN(n int) int {
	ns := strconv.Itoa(n)
	m := make([]int, 10)
	res := 0
	for i := 0; i < len(ns)-1; i++ {
		res += 9 * getChoices(i, 9)
	}
	cnt := 10
	// i表示前多少位相同
	for i := 0; i <= len(ns); i++ {
		if i == len(ns) {
			res++
		}
		c := int(ns[i] - '0')
		if m[int(ns[i-1]-'0')] == 1 {
			break
		}
		if i == 0 {
			c -= 1
		}
		cnt--
		if i > 0 {
			m[int(ns[i-1]-'0')] = 1
		}
		res += c * getChoices(len(ns)-i-1, cnt)
	}
	return n - res
}
func getChoices(l, c int) int {
	res := 1
	for i := 0; i < l; i++ {
		res *= c - i
	}
	return res
}
