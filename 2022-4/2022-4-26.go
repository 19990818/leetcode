package main

import (
	"strconv"
	"strings"
)

func shoppingOffers(price []int, special [][]int, needs []int) int {
	ans := 0
	//大礼包可以选 可以不选 假设最初为不选
	for i := 0; i < len(needs); i++ {
		ans += price[i] * needs[i]
	}
	var isOverflow func(needs []int, spec []int) bool
	isOverflow = func(needs, spec []int) bool {
		for i := 0; i < len(needs); i++ {
			if needs[i]-spec[i] < 0 {
				return true
			}
		}
		return false
	}
	var dfs func(needs []int, special [][]int, specialsum int)
	dfs = func(needs []int, special [][]int, specialsum int) {
		for _, val := range needs {
			if val < 0 {
				return
			}
		}
		for _, val := range special {
			temp := 0
			if isOverflow(needs, val) {
				continue
			}
			for i := 0; i < len(needs); i++ {
				needs[i] -= val[i]
			}
			specialsum += val[len(needs)]
			temp += specialsum
			for idx, need := range needs {
				temp += need * price[idx]
			}
			ans = min(ans, temp)
			dfs(needs, special, specialsum)
			for i := 0; i < len(needs); i++ {
				needs[i] += val[i]
			}
			specialsum -= val[len(needs)]
		}
	}
	dfs(needs, special, 0)
	return ans
}

func solveEquation(equation string) string {
	equalArr := strings.Split(equation, "=")
	left, right := equalArr[0], equalArr[1]
	var getXandDigit func(s string) (int, int)
	getXandDigit = func(s string) (int, int) {
		flag := 1
		countX, num := 0, 0
		for i := 0; i < len(s); i++ {
			if s[i] == '+' {
				flag = 1
			} else if s[i] == '-' {
				flag = -1
			} else if s[i] == 'x' {
				countX += flag
			} else {
				temp := ""
				for i < len(s) && s[i] <= '9' && s[i] >= '0' {
					temp += string(s[i])
					i++
				}
				tempNum, _ := strconv.Atoi(temp)
				if i < len(s) && s[i] == 'x' {
					countX += flag * tempNum
				} else {
					num += flag * tempNum
					i--
				}
			}
		}
		return countX, num
	}
	leftCountX, leftNum := getXandDigit(left)
	rightCountX, rightNum := getXandDigit(right)
	//fmt.Println(leftCountX,leftNum,rightCountX,rightNum)
	if leftCountX == rightCountX && leftNum == rightNum {
		return "Infinite solutions"
	}
	if leftCountX == rightCountX && leftNum != rightNum {
		return "No solution"
	}
	if (rightNum-leftNum)%(rightCountX-leftCountX) != 0 {
		return "No solution"
	}
	ans := (rightNum - leftNum) / (-rightCountX + leftCountX)
	var res strings.Builder
	ansS := strconv.Itoa(ans)
	res.WriteString("x=")
	res.WriteString(ansS)
	return res.String()
}

func countSubstrings(s string) int {
	dp := make([][]int, len(s))
	ans := len(s)
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		for j := 0; j <= i; j++ {
			dp[i][j] = 1
		}
	}
	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if dp[i+1][j-1] == 1 && s[i] == s[j] {
				//fmt.Println(i,j)
				dp[i][j] = 1
				ans++
			}
		}
	}
	return ans
}
