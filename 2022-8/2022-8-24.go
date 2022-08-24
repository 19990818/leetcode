package main

import (
	"strings"
)

func numBusesToDestination(routes [][]int, source int, target int) int {
	//换乘公交数 若使用bfs需要得到每个车站有哪些车次
	//然后我们通过车次路线可以转到下一个车站
	//busStation标记为一个车站所能通过的所有车次
	if source == target {
		return 0
	}
	busStation := make(map[int][]int)
	for i, val := range routes {
		for _, val2 := range val {
			busStation[val2] = append(busStation[val2], i)
		}
	}
	busM := make(map[int]int)
	cnt := 0
	q := []int{source}
	for {
		temp := make([]int, 0)
		for len(q) > 0 {
			cur := q[0]
			if cur == target {
				return cnt
			}
			q = q[1:]
			for _, bus := range busStation[cur] {
				if busM[bus] == 0 {
					busM[bus] = 1
					temp = append(temp, routes[bus]...)
				}
			}
		}
		if len(temp) == 0 {
			break
		}
		q = temp
		cnt++
	}
	return -1
}

func ambiguousCoordinates(s string) []string {
	n := len(s)
	//逗号必定在1~n-1中产生
	var isValid func(src string) bool
	isValid = func(src string) bool {
		if src == "" || src == "." {
			return false
		}
		arr := strings.Split(src, ".")
		if len(arr[0]) == 0 || (arr[0][0] == '0' && len(arr[0]) > 1) {
			return false
		}
		if len(arr) > 1 && (len(arr[1]) == 0 || arr[1][len(arr[1])-1] == '0') {
			return false
		}
		return true
	}
	m := make(map[string]int)
	ans := make([]string, 0)
	for i := 2; i < n-1; i++ {
		//fmt.Println("what", s[1:i], s[i:n-1])
		if isValid(s[1:i]) && isValid(s[i:n-1]) {
			ans = append(ans, s[0:i]+", "+s[i:])
		}
		for j := 1; j <= i; j++ {
			for k := i; k <= n-1; k++ {
				//fmt.Println("what")
				left, right := s[1:j]+"."+s[j:i], s[i:k]+"."+s[k:n-1]
				temp := ""
				if isValid(left) && isValid(right) {
					temp = "(" + left + ", " + right + ")"
				} else if isValid(left) && isValid(s[i:n-1]) {
					temp = "(" + left + ", " + s[i:]
				} else if isValid(right) && isValid(s[1:i]) {
					temp = s[0:i] + ", " + right + ")"
				}
				if temp != "" && m[temp] == 0 {
					ans = append(ans, temp)
					m[temp] = 1
				}
			}
		}
	}
	return ans
}

// "(123)"
//  "(00011)"
//  "(0123)"
