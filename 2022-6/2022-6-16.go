package main

import "strings"

func largestRectangleArea(heights []int) int {
	//动态规划解决失败，存在12345和找到的子问题无关
	//单调栈 当前元素小于栈顶元素 以当前元素结束的最大面积
	//每一个高度都可以作为矩形的长，求的每个高度的宽最长
	//可以得到每个高度的最大面积，然后再求最大值
	n := len(heights)
	if n == 1 {
		return heights[0]
	}
	var getConMaxLen func(start int, inc int) []int
	getConMaxLen = func(start int, inc int) []int {
		count := make([]int, n)
		stack := make([]int, 0)
		stack = append(stack, 0)
		i := start
		for ; i >= 0 && i < n; i = i + inc {
			for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
				count[stack[len(stack)-1]] = (i - stack[len(stack)-1]) * inc
				stack = stack[0 : len(stack)-1]
			}
			stack = append(stack, i)
		}
		for len(stack) > 0 {
			count[stack[len(stack)-1]] = (i - stack[len(stack)-1]) * inc
			stack = stack[0 : len(stack)-1]
		}
		return count
	}
	count := make([]int, n)
	countLeft := getConMaxLen(0, 1)
	countRight := getConMaxLen(n-1, -1)
	//fmt.Println(countLeft,countRight)
	for i := range count {
		count[i] = countLeft[i] + countRight[i] - 1
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, count[i]*heights[i])
	}
	return ans
}

func minWindow(s string, t string) string {
	//通过map得到t对应的元素以及数量
	m := make(map[byte]int)
	for i := range t {
		m[t[i]]++
	}
	totalCount := len(m)
	count := 0
	j := 0
	ans := ""
	for i := 0; i < len(s); i++ {
		for j < len(s) && count < totalCount {
			if _, ok := m[s[j]]; ok {
				m[s[j]]--
				if m[s[j]] == 0 {
					count++
				}
			}
			j++
		}
		if count >= totalCount {
			if ans == "" {
				ans = s[i:j]
			} else if len(ans) > j-i {
				ans = s[i:j]
			}
		}
		if _, ok := m[s[i]]; ok {
			m[s[i]]++
			if m[s[i]] > 0 {
				count--
			}
		}
	}
	return ans
}

func fullJustify(words []string, maxWidth int) []string {
	//每一行最大maxWidth 单词不可拆分 剩下的用空格均匀代替 若不能均匀分配 左边分配更多 %
	//若为n个单词 则其中分割为n-1
	pre := 0
	sum := 0
	ans := make([]string, 0)
	for idx := 0; idx <= len(words); idx++ {
		if idx < len(words) && idx != pre {
			sum += 1
		}
		if idx < len(words) {
			sum += len(words[idx])
		}
		if sum > maxWidth || idx == len(words) {
			//len=idx-pre
			temp := make([]byte, 0)
			if idx < len(words) {
				sum -= len(words[idx]) + 1
			}
			if idx-pre-1 == 0 {
				temp = append(temp, []byte(words[pre])...)
				for i := len(words[pre]); i < maxWidth; i++ {
					temp = append(temp, ' ')
				}
			} else {
				rest := (maxWidth - sum) / (idx - pre - 1)
				more := (maxWidth - sum) % (idx - pre - 1)
				for i := pre; i < idx; i++ {
					if i > pre {
						if i-pre <= more {
							temp = append(temp, ' ')
						}
						for j := 0; j < rest+1; j++ {
							temp = append(temp, ' ')
						}
					}
					temp = append(temp, []byte(words[i])...)
				}
			}
			ans = append(ans, string(temp))
			if idx < len(words) {
				sum = len(words[idx])
			}
			pre = idx
		}
	}
	//fmt.Println(ans[len(ans)-1])
	lastArr := strings.Split(ans[len(ans)-1], " ")
	var lastModify strings.Builder
	for _, val := range lastArr {
		if val == "" {
			continue
		}
		//fmt.Println(val)
		temp := strings.Trim(val, " ")
		//fmt.Println(temp)
		if lastModify.Len() != 0 {
			lastModify.WriteByte(' ')
		}
		lastModify.WriteString(temp)
	}
	for lastModify.Len() < maxWidth {
		lastModify.WriteString(" ")
	}
	ans[len(ans)-1] = lastModify.String()
	return ans
}
