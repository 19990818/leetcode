package main

import (
	"strconv"
	"strings"
)

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[i][j] = -1
		}
	}
	total := 0
	queue := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				ans[i][j] = 0
				queue = append(queue, []int{i, j})
				total++
			}

		}
	}
	count := 1
	for {
		temp := make([][]int, 0)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			travel := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			for _, val := range travel {
				if cur[0]+val[0] >= 0 && cur[0]+val[0] < m && cur[1]+val[1] >= 0 && cur[1]+val[1] < n && ans[cur[0]+val[0]][cur[1]+val[1]] == -1 {
					ans[cur[0]+val[0]][cur[1]+val[1]] = count
					total++
					temp = append(temp, []int{cur[0] + val[0], cur[1] + val[1]})
				}
			}
		}
		count++
		//fmt.Println(total,ans)
		if total >= m*n {
			break
		}
		queue = temp
	}
	return ans
}

func findCircleNum(isConnected [][]int) int {
	connect := make([][]int, len(isConnected))
	for index, val := range isConnected {
		for index2, val2 := range val {
			if val2 == 1 && index != index2 {
				connect[index] = append(connect[index], index2)
			}
		}
	}
	//fmt.Println(connect)
	flag := make([]int, len(isConnected))
	var dfs func(connect [][]int, index int)
	dfs = func(connect [][]int, index int) {
		queue := make([]int, 0)
		queue = append(queue, index)
		//fmt.Println(flag)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, val := range connect[cur] {
				if flag[val] == 0 {
					flag[val] = 1
					queue = append(queue, val)
				}
			}
			//fmt.Println("queue",queue)
		}
	}
	ans := 0
	for i := 0; i < len(isConnected); i++ {
		if flag[i] == 0 {
			flag[i] = 1
			ans++
			//fmt.Println(i)
			dfs(connect, i)
		}
	}
	return ans
}

func optimalDivision(nums []int) string {
	//最大 等于当前除以最小
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	var ans strings.Builder
	if len(nums) == 2 {
		ans.WriteString(strconv.Itoa(nums[0]))
		ans.WriteString("/")
		ans.WriteString(strconv.Itoa(nums[1]))
		return ans.String()
	}
	ans.WriteString(strconv.Itoa(nums[0]))
	ans.WriteString("/")
	ans.WriteString("(")
	ans.WriteString(getMinDiv(nums[1:]))
	ans.WriteString(")")
	return ans.String()
}
func getMinDiv(nums []int) string {
	var ans strings.Builder
	for idx, val := range nums {
		if idx > 0 {
			ans.WriteString("/")
		}
		ans.WriteString(strconv.Itoa(val))
	}
	return ans.String()
}
