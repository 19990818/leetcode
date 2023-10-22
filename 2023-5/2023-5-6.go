package main

func minNumberOfFrogs(croakOfFrogs string) int {
	m := make([]int, 5)
	res := 0
	cnt := 0
	h := map[rune]int{'c': 0, 'r': 1, 'o': 2, 'a': 3, 'k': 4}
	for _, v := range croakOfFrogs {
		if h[v] > 0 {
			if m[h[v]-1] <= 0 {
				return -1
			}
			m[h[v]-1]--
		} else {
			if cnt > 0 {
				cnt--
			} else {
				res++
			}

		}
		if h[v] == 4 {
			cnt++
		}
		m[h[v]]++
	}
	if checka(m) {
		return res
	}
	return -1
}
func checka(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] != 0 {
			return false
		}
	}
	return true
}
