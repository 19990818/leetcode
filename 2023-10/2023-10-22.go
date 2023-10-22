package main

import (
	"sort"
	"strconv"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	sum := 0
	res := 0
	for i := len(satisfaction) - 1; i >= 0 && sum >= 0; i-- {
		sum += satisfaction[i]
		if sum >= 0 {
			res += sum
		}
	}
	return res
}

func minProcessingTime(processorTime []int, tasks []int) int {
	sort.Ints(processorTime)
	sort.Ints(tasks)
	res, j := 0, 0
	for i := len(tasks) - 1; i >= 0; i -= 4 {
		res = max(res, tasks[i]+processorTime[j])
		j++
	}
	return res
}

func lastVisitedIntegers(words []string) []int {
	res, s := make([]int, 0), make([]string, 0)
	cnt := 0
	for _, v := range words {
		if v == "prev" {
			cnt++
			if cnt > len(s) {
				res = append(res, -1)
			} else {
				num, _ := strconv.Atoi(s[len(s)-cnt])
				res = append(res, num)
			}
		} else {
			cnt = 0
			s = append(s, v)
		}
	}
	return res
}

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	s := make([]int, 0)
	s = append(s, 0)
	for i := 1; i < len(groups); i++ {
		if groups[i] != groups[s[len(s)-1]] {
			s = append(s, i)
		}
	}
	res := make([]string, len(s))
	for i, v := range s {
		res[i] = words[v]
	}
	return res
}

func getWordsInLongestSubsequence2(n int, words []string, groups []int) []string {
	m := make(map[string][]int)
	m[words[0]] = []int{0}
	cnt := 1
	for i := 1; i < len(words); i++ {
		for _, v := range getNext(words[i]) {
			if len(m[v]) > 0 && groups[m[v][len(m[v])-1]] != groups[i] && len(m[v]) > len(m[words[i]]) {
				m[words[i]] = append([]int{}, m[v]...)
			}
		}
		m[words[i]] = append(m[words[i]], i)
		cnt = max(cnt, len(m[words[i]]))
	}
	res := make([]string, 0)
	for _, v := range m {
		if len(v) == cnt {
			for _, idx := range v {
				res = append(res, words[idx])
			}
			break
		}
	}
	return res
}
func getNext(s string) []string {
	res := make([]string, 0)
	temp := []byte(s)
	for i := range s {
		for j := 0; j < 26; j++ {
			if s[i] != byte(j+'a') {
				temp[i] = byte(j + 'a')
				res = append(res, string(temp))
			}
		}
		temp[i] = s[i]
	}
	return res
}
