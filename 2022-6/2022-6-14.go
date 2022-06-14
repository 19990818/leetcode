package main

import "strings"

func getPermutation(n int, k int) string {
	m := make(map[int]int)
	pool := make([]int, n)
	for i := range pool {
		pool[i] = i + 1
	}
	ans := make([]byte, n)
	//确定第一个元素
	var factor func(n int) int
	factor = func(n int) int {
		ans := 1
		for n > 0 {
			ans *= n
			n--
		}
		return ans
	}
	k = k - 1
	for i := 0; i < n; i++ {
		parts := factor(n - i - 1)
		pos := k / parts
		cnt := 0
		for _, val := range pool {
			if m[val] == 0 {
				if cnt == pos {
					m[val] = 1
					ans[i] = byte(val + '0')
					break
				}
				cnt++
			}
		}
		k = k % parts
	}
	return string(ans)
}

func isNumber(s string) bool {
	tempStr := strings.Split(s, "e")
	if len(tempStr) == 1 {
		tempStr = strings.Split(s, "E")
	}
	if len(tempStr) > 2 {
		return false
	}
	if len(tempStr) == 2 {
		if len(tempStr[0]) == 0 || len(tempStr[1]) == 0 {
			return false
		}
	}
	flag := false
	for j := 0; j < len(tempStr); j++ {
		count := j
		for i := 0; i < len(tempStr[j]); i++ {
			if i == 0 && (tempStr[j][0] == '+' || tempStr[j][0] == '-') {
				if len(tempStr[j]) == 1 {
					return false
				}
				continue
			}
			if tempStr[j][i] == '.' {
				count++
				if count > 1-j {
					return false
				}
			} else if tempStr[j][i] < '0' || tempStr[j][i] > '9' {
				return false
			} else {
				if j == 0 {
					flag = true
				}
			}
		}
	}
	return flag
}
