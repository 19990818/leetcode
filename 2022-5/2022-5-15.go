package main

import (
	"reflect"
	"sort"
)

func removeAnagrams(words []string) []string {
	var same func(a, b string) bool
	same = func(a, b string) bool {
		if len(a) != len(b) {
			return false
		}
		ma, mb := make(map[byte]int), make(map[byte]int)
		for i := 0; i < len(a); i++ {
			ma[a[i]]++
			mb[b[i]]++
		}
		return reflect.DeepEqual(ma, mb)
	}

	for {
		temp := make([]string, 0)
		for i := 1; i < len(words); i++ {
			if same(words[i], words[i-1]) {
				temp = append(temp, words[0:i]...)
				temp = append(temp, words[i+1:]...)
				break
			}
		}
		if len(temp) == len(words) {
			break
		}
		words = temp
	}
	return words
}

func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)
	cur := bottom
	ans := 0
	for _, val := range special {
		ans = max(ans, val-cur)
		cur = val + 1
	}
	ans = max(ans, top-cur+1)
	return ans
}

func largestCombination(candidates []int) int {
	var oneCounts func(a int) []int
	oneCounts = func(a int) []int {
		res := make([]int, 32)
		for i := 31; i >= 0; i-- {
			if a >= 1<<i {
				a -= 1 << i
				res[i] = 1
			} else {
				res[i] = 0
			}
		}
		return res
	}
	temp := make([][]int, 0)
	for _, val := range candidates {
		temp = append(temp, oneCounts(val))
	}
	ans := 0
	for j := 0; j < 32; j++ {
		count := 0
		for i := 0; i < len(candidates); i++ {
			if temp[i][j] == 1 {
				count++
			}
		}
		ans = max(ans, count)
	}
	return ans
}
