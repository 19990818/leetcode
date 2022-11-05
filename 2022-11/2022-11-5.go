package main

import (
	"strings"
)

func parseBoolExpr(expression string) bool {
	// 表达式只可能是以t,f开头 或者运算符
	//fmt.Println(expression)
	m := map[string]bool{"t": true, "f": false}
	if expression == "t" || expression == "f" {
		return m[expression]
	}
	sub := splitEpr(expression[2 : len(expression)-1])
	// fmt.Println(sub)
	res := parseBoolExpr(sub[0])
	switch expression[0] {
	case '&':
		for _, s := range sub {
			res = res && parseBoolExpr(s)
		}
	case '|':
		for _, s := range sub {
			res = res || parseBoolExpr(s)
		}
	case '!':
		for _, s := range sub {
			res = !parseBoolExpr(s)
		}
	}
	return res
}
func splitEpr(a string) []string {
	ans := make([]string, 0)
	//fmt.Println(a)
	//isOp := map[byte]int{'&': 1, '|': 1, '!': 1}
	var res strings.Builder
	stack := make([]byte, 0)
	// 加上一个尾巴方便统一处理
	a = a + ","
	for i := 0; i < len(a); i++ {
		if a[i] == '(' {
			stack = append(stack, a[i])
		} else if a[i] == ')' {
			stack = stack[0 : len(stack)-1]
		} else if len(stack) == 0 && a[i] == ',' {
			ans = append(ans, res.String())
			res.Reset()
			continue
		}
		res.WriteByte(a[i])
	}
	return ans
}

func stoneGameII(piles []int) int {
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	sum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sum[i] = sum[i+1] + piles[i]
	}
	for i := 1; i <= n; i++ {
		dp[n-1][i] = piles[n-1]
	}
	for i := n - 2; i >= 0; i-- {
		for j := 1; j <= n; j++ {
			for x := 0; i+x < n; x++ {
				if x >= 1 && x <= j {
					dp[i][j] = max(dp[i][j], sum[i]-dp[i+x][j])
				} else if x > j && x <= 2*j {
					dp[i][j] = max(dp[i][j], sum[i]-dp[i+x][x])
				}
			}
		}
	}
	return dp[0][1]
}

func powerfulIntegers(x int, y int, bound int) []int {
	arrx, arry := []int{1}, []int{1}
	if x != 1 {
		arrx = getSmallerBound(x, bound)
	}
	if y != 1 {
		arry = getSmallerBound(y, bound)
	}
	m := make(map[int]int)
	res := make([]int, 0)
	for i := range arrx {
		for j := range arry {
			if m[arrx[i]+arry[j]] == 0 && arrx[i]+arry[j] <= bound {
				res = append(res, arrx[i]+arry[j])
				m[arrx[i]+arry[j]] = 1
			}
		}
	}
	return res
}
func getSmallerBound(t, bound int) []int {
	temp := 1
	ans := []int{}
	for temp <= bound {
		ans = append(ans, temp)
		temp *= t
	}
	return ans
}
