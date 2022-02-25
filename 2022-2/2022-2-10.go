package main

import (
	"strconv"
	"strings"
)

func searchMatrix2(matrix [][]int, target int) bool {
	type cordinate struct {
		x int
		y int
	}
	cur := make([]cordinate, 0)
	cur = append(cur, cordinate{0, 0})
	if matrix[0][0] == target {
		return true
	}
	flag := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		flag[i] = make([]int, len(matrix[0]))
	}
	flag[0][0] = 1
	for len(cur) != 0 {
		if cur[0].x+1 < len(matrix) && matrix[cur[0].x+1][cur[0].y] == target {
			return true
		}
		if cur[0].x+1 < len(matrix) && matrix[cur[0].x+1][cur[0].y] < target {
			if flag[cur[0].x+1][cur[0].y] == 0 {
				cur = append(cur, cordinate{cur[0].x + 1, cur[0].y})
				flag[cur[0].x+1][cur[0].y] = 1
			}
		}
		if cur[0].y+1 < len(matrix[0]) && matrix[cur[0].x][cur[0].y+1] == target {
			return true
		}
		if cur[0].y+1 < len(matrix[0]) && matrix[cur[0].x][cur[0].y+1] < target {
			if flag[cur[0].x][cur[0].y+1] == 0 {
				cur = append(cur, cordinate{cur[0].x, cur[0].y + 1})
				flag[cur[0].x][cur[0].y+1] = 1
			}
		}
		cur = cur[1:]
	}
	return false
}

func searchMatrix3(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

func diffWaysToCompute(expression string) []int {
	// 是否只有两种选择a和b在一起 a和b不在一起
	if isDigit(expression) {
		res, _ := strconv.Atoi(expression)
		return []int{res}
	}
	ans := make([]int, 0)
	for idx, val := range expression {
		if val == '+' || val == '-' || val == '*' {
			leftNums := diffWaysToCompute(expression[0:idx])
			rightNums := diffWaysToCompute(expression[idx+1:])
			for _, leftNum := range leftNums {
				for _, rightNum := range rightNums {
					temp := 0
					if val == '+' {
						temp = leftNum + rightNum
					} else if val == '-' {
						temp = leftNum - rightNum
					} else {
						temp = leftNum * rightNum
					}
					ans = append(ans, temp)
				}
			}
		}
	}
	return ans
}
func isDigit(expression string) bool {
	_, err := strconv.Atoi(expression)
	if err != nil {
		return false
	}
	return true
}

func uncommonFromSentences(s1 string, s2 string) []string {
	m := make(map[string]int)
	s1Arr := strings.Split(s1, " ")
	s2Arr := strings.Split(s2, " ")
	ans := make([]string, 0)
	for _, val := range s1Arr {
		m[val]++
	}
	for _, val := range s2Arr {
		m[val]++
	}
	for key, val := range m {
		if val == 1 {
			ans = append(ans, key)
		}
	}
	return ans
}
