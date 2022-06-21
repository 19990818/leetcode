package main

import (
	"strconv"
	"strings"
)

func maxPathSum(root *TreeNode) int {
	//每个节点都会有一个右边的最大值，以及左边的最大值
	//然后经过此节点的路径最大值为maxright+maxleft-node.val
	var maxRight func(root *TreeNode) int
	var maxleft func(root *TreeNode) int
	mR, mL := make(map[*TreeNode]int), make(map[*TreeNode]int)
	maxRight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if _, ok := mR[root]; ok {
			return mR[root]
		}
		mR[root] = max(root.Val, root.Val+max(maxRight(root.Right), maxleft(root.Right)))
		return mR[root]
	}
	maxleft = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if _, ok := mL[root]; ok {
			return mL[root]
		}
		mL[root] = max(root.Val, root.Val+max(maxRight(root.Left), maxleft(root.Left)))
		return mL[root]
	}
	var dfs func(root *TreeNode)
	ans := root.Val
	dfs = func(root *TreeNode) {
		ans = max(ans, maxleft(root)+maxRight(root)-root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return ans
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	//这题关键是如何构造图，然后使用bfs即可
	//构建图时候可以借助一个虚拟的中间状态来表示两者之间可以改变一个状态得到
	//比如aba=>a*a=>aca 而不是通过遍历改变每个字符在列表中是否存在
	//引进中间状态实际上是用两个之间的正则表达式，比如a*a能表示aba与aca
	m := make(map[string][]string)
	wordList = append(wordList, beginWord)
	for _, val := range wordList {
		for i := 0; i < len(val); i++ {
			var temp strings.Builder
			temp.WriteString(val[0:i])
			temp.WriteString("*")
			temp.WriteString(val[i+1:])
			m[temp.String()] = append(m[temp.String()], val)
			m[val] = append(m[val], temp.String())
		}
	}
	if _, ok := m[endWord]; !ok {
		return 0
	}
	travel := make(map[string]int)
	travel[beginWord] = 1
	res := 0
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	for {
		temp := make([]string, 0)
		res += 1
		//fmt.Println(queue)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			travel[cur] = 1
			for _, next := range m[cur] {
				if next == endWord {
					return res/2 + 1
				}
				if _, ok := travel[next]; !ok {
					temp = append(temp, next)
				}
			}
		}
		if len(temp) == 0 {
			break
		}
		queue = temp
	}
	return 0
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	m := make(map[string][]string)
	wordList = append(wordList, beginWord)
	for _, val := range wordList {
		for i := 0; i < len(val); i++ {
			var temp strings.Builder
			temp.WriteString(val[0:i])
			temp.WriteString("*")
			temp.WriteString(val[i+1:])
			m[temp.String()] = append(m[temp.String()], val)
			m[val] = append(m[val], temp.String())
		}
	}
	if _, ok := m[endWord]; !ok {
		return [][]string{}
	}
	travel := make(map[string]int)
	res := make([][]string, 0)
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	var bfs func(queue []string) int
	bfs = func(queue []string) int {
		res := 0
		for {
			temp := make([]string, 0)
			res += 1
			//fmt.Println(queue)
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				travel[cur] = 1
				for _, next := range m[cur] {
					if next == endWord {
						return res
					}
					if _, ok := travel[next]; !ok {
						temp = append(temp, next)
					}
				}
			}
			if len(temp) == 0 {
				break
			}
			queue = temp
		}
		return res
	}
	maxSteps := bfs(queue)
	temp := make([]string, 0)
	var dfs func(cur string, steps int)
	dfs = func(cur string, steps int) {
		if steps > maxSteps {
			return
		}
		if cur == endWord {
			temp2 := append([]string{beginWord}, temp...)
			res = append(res, temp2)
		}
		for _, val := range m[cur] {
			if !strings.Contains(cur, "*") {
				temp = append(temp, val)
			}

			dfs(val, steps+1)
			if !strings.Contains(cur, "*") {
				temp = temp[0 : len(temp)-1]
			}

		}
	}
	dfs(beginWord, 0)
	return res
}

func findEvenNumbers(digits []int) []int {
	m := make(map[int]int)
	for _, val := range digits {
		m[val]++
	}
	ans := make([]int, 0)
	for i := 100; i < 999; i += 2 {
		str := strconv.Itoa(i)
		m2 := make(map[int]int)
		for _, val := range str {
			m2[int(val-'0')]++
		}
		flag := 0
		for key, val := range m2 {
			if m[key] < val {
				flag = 1
				break
			}
		}
		if flag == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}
