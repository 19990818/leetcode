package main

import (
	"strconv"
	"strings"
)

func calculate2(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	stack := make([]string, 0)
	temp := ""
	i := 0
	for i < len(s) {
		for i < len(s) && s[i] <= '9' && s[i] >= '0' {
			temp += string(s[i])
			i++
		}
		if i >= len(s) {
			stack = append(stack, temp)
			break
		}
		if s[i] == '*' || s[i] == '/' {
			if temp != "" {
				stack = append(stack, temp)
			}
			temp = ""
			j := i + 1
			for j < len(s) && s[j] <= '9' && s[j] >= '0' {
				temp += string(s[j])
				j++
			}
			a, _ := strconv.Atoi(stack[len(stack)-1])
			b, _ := strconv.Atoi(temp)
			//fmt.Println(a,b,stack)
			if s[i] == '*' {
				a = a * b
				stack[len(stack)-1] = strconv.Itoa(a)
			} else {
				a = a / b
				stack[len(stack)-1] = strconv.Itoa(a)
			}
			i = j
			temp = ""
		} else {
			stack = append(stack, temp)
			temp = ""
			stack = append(stack, string(s[i]))
			i++
		}
	}
	// fmt.Println(stack)
	ans, _ := strconv.Atoi(stack[0])
	for i := 1; i < len(stack); i++ {
		if stack[i] == "+" {
			temp, _ := strconv.Atoi(stack[i+1])
			ans += temp
		} else if stack[i] == "-" {
			temp, _ := strconv.Atoi(stack[i+1])
			ans -= temp
		}
	}
	return ans
}

func kthSmallest(root *TreeNode, k int) int {
	res := dfsPre(root)
	return res[k-1]
}
func dfsPre(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = dfsPre(root.Left)
	res = append(res, root.Val)
	res = append(res, dfsPre(root.Right)...)
	return res
}
