package main

import "strings"

func candy(ratings []int) int {
	n := len(ratings)
	grade := make([]int, n)
	for i := range grade {
		grade[i] = 1
	}
	var travelOneSide func(start, flag int) []int
	travelOneSide = func(start, flag int) []int {
		stack := make([]int, 0)
		stack = append(stack, start)
		res := make([]int, n)
		for i := start + flag; i >= 0 && i < n; i += flag {
			//fmt.Println(flag, stack)
			if len(stack) == 0 || ratings[i] < ratings[stack[len(stack)-1]] {
				stack = append(stack, i)
			} else {
				count := 0
				for len(stack) > 0 {
					cur := stack[len(stack)-1]
					res[cur] = count
					count++
					stack = stack[0 : len(stack)-1]
				}
				stack = append(stack, i)
			}
		}
		count := 0
		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			res[cur] = count
			count++
			stack = stack[0 : len(stack)-1]
		}
		return res
	}
	ans := n
	leftInc := travelOneSide(0, 1)
	RightInc := travelOneSide(n-1, -1)
	for i := 0; i < n; i++ {
		ans += max(leftInc[i], RightInc[i])
	}
	return ans
}

func wordBreak(s string, wordDict []string) []string {
	m := make(map[string][]string)
	wordM := make(map[string]int)
	for _, val := range wordDict {
		wordM[val] = 1
	}
	var recur func(s string) []string
	recur = func(s string) []string {
		if _, ok := m[s]; ok {
			return m[s]
		}
		res := make([]string, 0)
		for i := 1; i <= len(s); i++ {
			if _, ok := wordM[s[0:i]]; ok {
				if i == len(s) {
					res = append(res, s[0:i])
				} else {
					for _, val := range recur(s[i:]) {
						var temp strings.Builder
						temp.WriteString(s[0:i])
						temp.WriteString(" ")
						temp.WriteString(val)
						res = append(res, temp.String())
					}
				}
			}
		}
		m[s] = res
		return res
	}
	return recur(s)
}

func minCut(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	//dp[i][j]表示以i开始以j结束的字符串是否为回文
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				if i+1 > j-1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
		}
	}
	graph := make([][]int, n)
	for i := range dp {
		for j := range dp[i] {
			if dp[i][j] {
				graph[i] = append(graph[i], j+1)
			}
		}
	}
	//fmt.Println(graph)
	//fmt.Println("test")
	queue := make([]int, 0)
	queue = append(queue, 0)
	travel := make(map[int]int)
	ans := 0
	for {
		temp := make([]int, 0)
		for len(queue) > 0 {
			cur := queue[0]
			travel[cur] = 1
			queue = queue[1:]
			if cur == n {
				return ans - 1
			}
			for _, val := range graph[cur] {
				if travel[val] == 0 {
					travel[val] = 1
					temp = append(temp, val)
				}
			}
			//fmt.Println(queue)
		}
		ans++
		//fmt.Println(temp)
		if len(temp) == 0 {
			break
		}
		queue = temp
	}
	return 0
}
