package main

func longestWPI(hours []int) int {
	m := make(map[int]int)
	m[0] = -1
	cnt := 0
	res := 0
	for i, v := range hours {
		if v > 8 {
			cnt++
		} else {
			cnt--
		}
		if _, ok := m[cnt]; !ok {
			m[cnt] = i
		}
		if cnt > 0 {
			res = max(res, i+1)
		} else {
			if _, ok := m[cnt-1]; ok {
				res = max(res, i-m[cnt-1])
			}
		}
	}
	return res
}
