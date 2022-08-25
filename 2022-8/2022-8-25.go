package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	lab := []byte{}
	type a struct {
		index  int
		src    string
		target string
	}
	as := make([]a, 0)
	for i := range indices {
		as = append(as, a{indices[i], sources[i], targets[i]})
	}
	sort.Slice(as, func(i, j int) bool {
		return as[i].index < as[j].index
	})
	pre := 0
	for i := range as {
		indice := as[i].index
		if pre < indice {
			lab = append(lab, s[pre:indice]...)
			pre = indice
		}

		if indice+len(as[i].src) <= len(s) && s[indice:indice+len(as[i].src)] == as[i].src {
			lab = append(lab, as[i].target...)
			pre = indice + len(as[i].src)
		}

	}
	lab = append(lab, s[pre:]...)
	return string(lab)
}

func canVisitAllRooms(rooms [][]int) bool {
	queue := []int{0}
	m := make(map[int]int)
	m[0] = 1
	for len(queue) > 0 {
		cur := queue[0]
		m[cur] = 1
		for _, val := range rooms[cur] {
			if m[val] == 0 {
				m[val] = 1
				queue = append(queue, val)
			}
		}
	}
	return len(m) == len(rooms)
}

func splitIntoFibonacci(num string) []int {
	if len(num) < 3 {
		return []int{}
	}
	//确定两个开始的数值
	n := len(num)
	start1, start2 := 0, 0
	for i := 0; i <= n/2; i++ {
		if num[0] != '0' {
			start1, _ = strconv.Atoi(num[0 : i+1])
		} else if num[0] == '0' && i > 0 {
			break
		}
		j := i + 1
		start2 = 0
		for ; j < n-1; j++ {
			if num[i+1] != '0' {
				start2, _ = strconv.Atoi(num[i+1 : j+1])
			} else if num[i+1] == '0' && j > i+1 {
				break
			}
			res := make([]int, 0)
			var temp strings.Builder
			//fmt.Println(start1,start2)
			res = []int{start1, start2}
			temp.WriteString(num[0 : j+1])
			pre1, pre2 := start1, start2
			for temp.Len() < len(num) {
				cur := pre1 + pre2
				if cur > math.MaxInt32 {
					break
				}
				pre1, pre2 = pre2, cur
				res = append(res, cur)
				temp.WriteString(strconv.Itoa(cur))
				if temp.String() == num {
					return res
				}
			}
		}

	}
	return []int{}
}
