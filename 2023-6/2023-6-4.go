package main

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	m := make(map[float64]int)
	for i := 0; i < len(nums)/2; i++ {
		m[float64(nums[i]+nums[len(nums)-i-1])/float64(2)]++
	}
	return len(m)
}

func minimizedStringLength(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	return len(m)
}

func semiOrderedPermutation(nums []int) int {
	pos1, posn := 0, 0
	for i, v := range nums {
		if v == 1 {
			pos1 = i
		}
		if v == len(nums) {
			posn = i
		}
	}
	res := pos1 - 1 + len(nums) - posn
	if pos1 > posn {
		res -= 1
	}
	return res
}

func matrixSumQueries(n int, queries [][]int) int64 {
	colM := make(map[int][]int)
	rowM := make(map[int][]int)
	for i := len(queries) - 1; i >= 0; i-- {
		if queries[i][0] == 0 {
			if _, ok := rowM[queries[i][1]]; !ok {
				rowM[queries[i][1]] = []int{queries[i][2], len(colM)}
			}
		} else {
			if _, ok := colM[queries[i][1]]; !ok {
				colM[queries[i][1]] = []int{queries[i][2], len(rowM)}
			}
		}
	}
	//fmt.Println(colM,rowM)
	res := int64(0)
	for _, v := range rowM {
		res += int64((n - v[1]) * v[0])
	}
	for _, v := range colM {
		res += int64((n - v[1]) * v[0])
	}
	return res
}
