package main

import (
	"sort"
	"strings"
)

// #761 特殊的二进制序列
// 此题的关键是得到特殊的二进制序列必须以1开始以0结束
// 这样我们就可以将大问题变为小问题 从而解决
// 我们可以将这样的序列进行排序就可以得到答案 因为特殊序列会作为一个整体
// 每段之间并没有联系 所以可以区域和最大得到总和最大
func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	cnt := 0
	left := 0
	ans := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				temp := "1" + makeLargestSpecial(s[left+1:i]) + "0"
				ans = append(ans, temp)
				left = i + 1
			}
		}
	}
	var res strings.Builder
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	for _, val := range ans {
		res.WriteString(val)
	}
	return res.String()
}
