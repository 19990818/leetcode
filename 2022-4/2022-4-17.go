package main

import (
	"strconv"
	"strings"
)

func digitSum(s string, k int) string {
	var ans strings.Builder
	ans.WriteString(s)
	for ans.Len() > 3 {
		var tempStr strings.Builder
		for i := 0; i < len(ans.String()); {
			temp := 0
			for j := i; j < len(ans.String()) && j < i+3; j++ {
				temp += int(ans.String()[j] - '0')
			}
			i += 3
			tempStr.WriteString(strconv.Itoa(temp))
		}
		ans.Reset()
		ans.WriteString(tempStr.String())
	}
	return ans.String()
}

func minimumRounds(tasks []int) int {
	m := make(map[int]int)
	for _, val := range tasks {
		m[val]++
	}
	ans := 0
	for _, val := range m {
		if val < 2 {
			return -1
		}
		flag := 0
		for i := val / 3; i >= 0; i-- {
			if (val-3*i)%2 == 0 {
				ans += (i + (val-3*i)/2)
				flag = 1
				break
			}
		}
		if flag == 0 {
			return -1
		}
	}
	return ans
}

func maxTrailingZeros(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	y, x := make([][][]int, m+1), make([][][]int, m+1)
	for i := 0; i <= m; i++ {
		x[i] = make([][]int, n+1)
		y[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			x[i][j] = make([]int, 2)
			y[i][j] = make([]int, 2)
		}
	}
	var getFiveAndTwo func(num int) []int
	getFiveAndTwo = func(num int) []int {
		five, two := 0, 0
		for num%5 == 0 {
			five++
			num /= 5
		}
		for num%2 == 0 {
			two++
			num /= 2
		}
		return []int{five, two}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp := getFiveAndTwo(grid[i][j])
			x[i+1][j+1][0] = x[i+1][j][0] + temp[0]
			x[i+1][j+1][1] = x[i+1][j][1] + temp[1]
			y[i+1][j+1][0] = y[i][j+1][0] + temp[0]
			y[i+1][j+1][1] = y[i][j+1][1] + temp[1]
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp := getFiveAndTwo(grid[i][j])
			t1 := min(y[i+1][j+1][0]+x[i+1][j+1][0]-temp[0], y[i+1][j+1][1]+x[i+1][j+1][1]-temp[1])
			t2 := min(y[m][j+1][0]-y[i+1][j+1][0]+x[i+1][j+1][0], y[m][j+1][1]-y[i+1][j+1][1]+x[i+1][j+1][1])
			t3 := min(y[i+1][j+1][0]+x[i+1][n][0]-x[i+1][j+1][0], y[i+1][j+1][1]+x[i+1][n][1]-x[i+1][j+1][1])
			t4 := min(y[m][j+1][0]-y[i+1][j+1][0]+x[i+1][n][0]-x[i+1][j+1][0]+temp[0], y[m][j+1][1]-y[i+1][j+1][1]+x[i+1][n][1]-x[i+1][j+1][1]+temp[1])
			com := []int{t1, t2, t3, t4}
			for _, val := range com {
				//fmt.Println(val,temp)
				ans = max(ans, val)
			}
		}
	}
	return ans
}
