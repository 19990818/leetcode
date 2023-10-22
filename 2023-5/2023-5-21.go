package main

import (
	"math"
	"strconv"
	"strings"
)

func storeWater(bucket []int, vat []int) int {
	maxAdd := 0
	zero := 0
	for i := range bucket {
		if bucket[i] == 0 {
			zero++
			bucket[i]++
		}
		maxAdd = max(maxAdd, vat[i]/bucket[i])
	}
	res := math.MaxInt64
	if maxAdd == 0 {
		return 0
	}
	for i := 1; i <= maxAdd+1; i++ {
		bigger := 0
		for j, v := range vat {
			bigger += max((v+i-1)/i-bucket[j], 0)
		}
		res = min(res, zero+bigger+i)
	}
	return res
}

func minLength(s string) int {
	t := ""
	for t != s {
		t = s
		s = strings.ReplaceAll(s, "AB", "")
		s = strings.ReplaceAll(s, "CD", "")
	}
	return len(s)
}

func makeSmallestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s)/2; i++ {
		if s[i] == s[len(s)-i-1] || s[i] < s[len(s)-i-1] {
			res += string(s[i])
		} else {
			res += string(s[len(s)-i-1])
		}
	}
	if len(s)%2 == 1 {
		res += string(s[len(s)/2])
	}
	for i := 0; i < len(s)/2; i++ {
		res += string(res[len(s)/2-i-1])
	}
	return res
}
func punishmentNumber(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		if checkpunish(i) {
			res += i * i
		}
	}
	return res
}
func checkpunish(i int) bool {
	if i == 1000 {
		return true
	}
	s := strconv.Itoa(i * i)
	l := (len(s) + 1) / 2
	for j := 0; j < len(s)-l+1; j++ {
		num, _ := strconv.Atoi((s[j : j+l]))
		for _, v1 := range getStr(s[0:j]) {
			for _, v2 := range getStr(s[j+l:]) {
				if num+v1+v2 == i {
					return true
				}
			}
		}
	}
	return false
}

// 得到一个字符串可能能够得到的值的组合
func getStr(a string) []int { 
	if len(a) > 3 || len(a) == 0 {
		return []int{0}
	}
	if len(a) == 1 {
		return []int{int(a[0] - '0')}
	}
	if len(a) == 2 {
		return []int{int(a[0]-'0') + int(a[1]-'0'),
			int(a[0]-'0')*10 + int(a[1]-'0')}
	}
	num1, _ := strconv.Atoi(a)
	num2, _ := strconv.Atoi(a[0:2])
	num3, _ := strconv.Atoi(a[1:3])
	return []int{num1, num2 + int(a[2]-'0'), num3 + int(a[0]-'0'),
		int(a[0]-'0') + int(a[1]-'0') + int(a[2]-'0')}
}
