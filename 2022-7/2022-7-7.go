package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func hjErrorRecord() {
	type flag struct {
		s    string
		line int
	}
	queue := make([]flag, 0)
	m := make(map[flag]int)
	for {
		var temp1 string
		var line int
		_, err := fmt.Scan(&temp1, &line)
		if err != nil {
			break
		}
		s := getFileName(temp1, 16)
		flagCur := flag{s, line}
		if m[flagCur] == 0 {
			queue = append(queue, flagCur)
		}
		m[flagCur]++
	}
	for i := max(len(queue)-8, 0); i < len(queue); i++ {
		fmt.Println(queue[i].s, queue[i].line, m[queue[i]])
	}
}
func getFileName(pathStr string, maxLen int) string {
	strArr := strings.Split(pathStr, "\\")
	lastStr := strArr[len(strArr)-1]
	if len(lastStr) > maxLen {
		return lastStr[len(lastStr)-maxLen:]
	}
	return lastStr
}

func main() {
	var n int
	fmt.Scan(&n)
	dic := make([]string, 0)
	for i := 0; i < n; i++ {
		var temp string
		fmt.Scan(&temp)
		dic = append(dic, temp)
	}
	var target string
	var k int
	fmt.Scan(&target, &k)
	res := getBrothers(target, dic)
	sort.Strings(res)
	fmt.Println(res)
	if k < len(res) {
		fmt.Println(res[k-1])
	}
}
func getBrothers(target string, dic []string) []string {
	var getM func(s string) map[rune]int
	getM = func(s string) map[rune]int {
		m := make(map[rune]int)
		for _, val := range s {
			m[val]++
		}
		return m
	}
	targetM := getM(target)
	ans := make([]string, 0)
	for _, val := range dic {
		if val != target {
			if reflect.DeepEqual(targetM, getM(val)) {
				ans = append(ans, val)
			}
		}
	}
	return ans
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	//并查集 帮派合并问题 当出现一个不存在在集合中的元素
	//是自己新建一个帮派还是加入帮派，这取决与equations关系
	//和之前的帮派有关系 就加入 没关系就新建
	//用数组来表示上下级关系 中间帮派之间有人有联系 撮合帮派合并
	//找到编号为n的上级 并查集关键的操作是find和merge
	//当新处理一条边的时候 若两者在同一集合中 使用find更新其帮派
	//不在同一集合中 将其中一个帮派的老大作为另一个帮派老大的下级
	//虽然这时候其中被合并的帮派的下级没有指向合并之后老大，但是
	//在使用的时候会再次进行find操作，会把这些没有指向合并之后老大的下级
	//全部指向老大
	n := len(equations)
	//表示每个变量的上级 最多2*n个变量
	parent := make([]int, 2*n)
	//权重初始都为1 方便乘积
	weight := make([]float64, 2*n)
	for i := 0; i < 2*n; i++ {
		parent[i] = -1
		weight[i] = 1.0
	}
	var find func(n int) int
	find = func(n int) int {
		//不存在上级 那么自己成为帮主
		if parent[n] < 0 {
			return n
		}
		//继续往上找
		father := find(parent[n])
		//注意在此处必须是乘以当前上级的权重值，father实际上为上上级
		weight[n] = weight[n] * weight[parent[n]]
		parent[n] = father
		return father
	}
	//帮派合并
	var merge func(a, b int, val float64)
	merge = func(a, b int, val float64) {
		a_root := find(a)
		b_root := find(b)
		if a_root == b_root {
			return
		}
		parent[a_root] = b_root
		//当存在四个点的时候 已知三条边的距离 根据乘积相当求出新建的边
		weight[a_root] = val * weight[b] / weight[a]
	}
	//针对不同的变量编号
	m := make(map[string]int)
	i := 0
	for idx, val := range equations {
		if _, ok := m[val[0]]; !ok {
			m[val[0]] = i
			i++
		}
		if _, ok := m[val[1]]; !ok {
			m[val[1]] = i
			i++
		}
		//新产生的直接进行并查集操作
		a_idx := m[val[0]]
		b_idx := m[val[1]]
		merge(a_idx, b_idx, values[idx])
	}
	ans := make([]float64, len(queries))
	for i, val := range queries {
		_, ok := m[val[0]]
		_, ok2 := m[val[1]]
		if !ok || !ok2 {
			ans[i] = -1.0
			continue
		}
		if find(m[val[0]]) != find(m[val[1]]) {
			ans[i] = -1.0
		} else {
			ans[i] = weight[m[val[0]]] / weight[m[val[1]]]
		}
	}
	return ans
}
