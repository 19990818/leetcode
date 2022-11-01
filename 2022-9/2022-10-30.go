package main

import (
	"math"
	"strconv"
)

func averageValue(nums []int) int {
	sum, cnt := 0, 0
	for _, num := range nums {
		if num%2 == 0 && num%3 == 0 {
			sum += num
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	mv := 0
	m := make(map[string]int)
	v := make(map[string][]int)
	n := len(creators)
	for i := 0; i < n; i++ {
		m[creators[i]] += views[i]
		v[creators[i]] = append(v[creators[i]], i)
	}
	for _, v := range m {
		mv = max(mv, v)
	}
	mc := make([]string, 0)
	for k, v := range m {
		if v == mv {
			mc = append(mc, k)
		}
	}
	ans := make([][]string, 0)
	for _, c := range mc {
		temp := views[v[c][0]]
		tempS := ids[v[c][0]]
		for i := 1; i < len(v[c]); i++ {
			if views[v[c][i]] > temp {
				temp = views[v[c][i]]
				tempS = ids[v[c][i]]
			} else if views[v[c][i]] == temp && ids[v[c][i]] < tempS {
				tempS = ids[v[c][i]]
			}
		}
		ans = append(ans, []string{c, tempS})
	}
	return ans
}

func makeIntegerBeautiful(n int64, target int) int64 {
	if n <= int64(target) {
		return 0
	}
	ns := strconv.FormatInt(n, 10)
	narr := make([]int, len(ns))
	sum := 0
	t := int64(0)
	end := -1
	for i, v := range ns {
		narr[i] = int(v - '0')
		sum += narr[i]
		if sum >= target {
			end = i
			if sum == target {
				flag := true
				for j := end + 1; j < len(ns); j++ {
					if ns[j] != '0' {
						flag = false
					}
				}
				if flag {
					return 0
				}
			}

			break
		}
	}
	//fmt.Println(end, sum)
	if end == -1 {
		return 0
	}
	if end == 0 {
		t = int64(math.Pow10(len(ns)))
	} else {
		for i := 0; i < end-1; i++ {
			t += int64(narr[i] * int(math.Pow10(len(ns)-i-1)))
		}
		t += int64(narr[end-1]) * int64(math.Pow10(len(ns)-end))
		t += int64(math.Pow10(len(ns) - end))
	}
	return t - n
}
