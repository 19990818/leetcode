package main

import (
	"strconv"
	"strings"
)

//在处理的过程中对题意理解不到位 导致记录之前的出现的不是使用栈
//而是和前置时间点的常量 导致有的任务结束后 持续的时间本应该是前一个任务
//而将此时间算到结束的任务当中
func exclusiveTime(n int, logs []string) []int {
	ans := make([]int, n)
	stack := make([]int, 0)
	strArr0 := strings.Split(logs[0], ":")
	num0, _ := strconv.Atoi(strArr0[0])
	num20, _ := strconv.Atoi(strArr0[2])
	stack = append(stack, num0)
	for i := 1; i < len(logs); i++ {
		strArri := strings.Split(logs[i], ":")
		numCur, _ := strconv.Atoi(strArri[0])
		num2Cur, _ := strconv.Atoi(strArri[2])
		if strArri[1] == "start" {
			if len(stack) != 0 {
				ans[stack[len(stack)-1]] += num2Cur - num20
			}
			num20 = num2Cur
			stack = append(stack, numCur)
		} else {
			ans[stack[len(stack)-1]] += num2Cur - num20 + 1
			stack = stack[0 : len(stack)-1]
			num20 = num2Cur + 1
		}
	}
	return ans
}
