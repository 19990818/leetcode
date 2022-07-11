package main

import (
	"sort"
	"strconv"
)

func lenLongestFibSubseq(arr []int) int {
	//统计以每个结束的最长斐波那契
	m := make(map[int]int)
	//得到值和下标的映射
	for i, val := range arr {
		m[val] = i
	}
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	ans := 0
	//分别以下标i,j为后两项的斐波拉契
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			target := arr[j] - arr[i]
			if _, ok := m[target]; ok && target < arr[i] {
				dp[i][j] = max(dp[i][j], dp[m[target]][i]+1)
			} else {
				dp[i][j] = 2
			}
			ans = max(ans, dp[i][j])
		}
	}
	if ans < 3 {
		return 0
	}
	return ans
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	//针对体重最大的人进行处理 双指针
	i, j := 0, len(people)-1
	ans := 0
	for i < j {
		//只能运走j
		if people[j]+people[i] > limit {
			j--
		} else {
			//运走i和j
			j--
			i++
		}
		ans++
	}
	//当这两者相等时 说明没有被运走
	if j == i {
		ans++
	}
	return ans
}

func scoreOfParentheses(s string) int {
	stack := make([]string, 0)
	for _, val := range s {
		if val == '(' {
			stack = append(stack, "(")
		} else {
			temp := 1
			in := 0
			for stack[len(stack)-1] != "(" {
				a, _ := strconv.Atoi(stack[len(stack)-1])
				stack = stack[0 : len(stack)-1]
				in += a
			}
			if in != 0 {
				temp = in * 2
			}
			//处理掉左括号
			stack = stack[0 : len(stack)-1]
			stack = append(stack, strconv.Itoa(temp))
		}
	}
	//fmt.Println(stack)
	ans := 0
	for len(stack) > 0 {
		temp, _ := strconv.Atoi(stack[0])
		//fmt.Println(temp)
		stack = stack[1:len(stack)]
		ans += temp
	}
	return ans
}
