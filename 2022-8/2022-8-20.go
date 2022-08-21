package main

import (
	"math"
	"sort"
	"strings"
)

func champagneTower(poured int, query_row int, query_glass int) float64 {
	// 简单动态规划
	dp := make([]float64, query_row+1)
	dp[0] = float64(poured)
	for i := 1; i <= query_row; i++ {
		temp := append([]float64{}, dp...)
		for j := 0; j < i+1; j++ {
			if j > 0 && j < i {
				dp[j] = (math.Max(temp[j-1]-1, 0) + math.Max(temp[j]-1, 0)) / 2
			} else if j == 0 {
				dp[j] = (math.Max(temp[j]-1, 0)) / 2
			} else {
				dp[j] = (math.Max(temp[j-1]-1, 0)) / 2
			}
		}
	}
	return math.Min(dp[query_glass], 1)
}

func eventualSafeNodes(graph [][]int) []int {
	// 实际上为拓扑排序找到安全节点
	n := len(graph)
	in, out := make([][]int, n), make([]int, n)
	queue := make([]int, 0)
	for idx, val := range graph {
		out[idx] = len(val)
		if out[idx] == 0 {
			queue = append(queue, idx)
		}
		for _, val2 := range val {
			in[val2] = append(in[val2], idx)
		}
	}
	ans := make([]int, 0)
	for len(queue) > 0 {
		cur := queue[0]
		ans = append(ans, cur)
		queue = queue[1:]
		for _, val := range in[cur] {
			out[val]--
			if out[val] == 0 {
				queue = append(queue, val)
			}
		}
	}
	sort.Ints(ans)
	return ans
}

func expressiveWords(s string, words []string) int {
	type pair struct {
		b   byte
		cnt int
	}
	var transferS func(s string) []pair
	transferS = func(s string) []pair {
		cur := s[0]
		cnt := 1
		res := make([]pair, 0)
		for i := 1; i <= len(s); i++ {
			if i == len(s) || s[i] != cur {
				res = append(res, pair{cur, cnt})
				cnt = 1
				if i < len(s) {
					cur = s[i]
				}

			} else {
				cnt++
			}
		}
		return res
	}
	target := transferS(s)
	ans := 0
	for _, val := range words {
		src := transferS(val)
		//fmt.Println(target,src)
		if len(src) == len(target) {
			ans++
			for idx := range src {
				if src[idx].b == target[idx].b && (src[idx].cnt == target[idx].cnt || target[idx].cnt >= 3 && target[idx].cnt > src[idx].cnt) {
					continue
				}
				ans--
				break
			}
		}
	}
	return ans
}

func minimumRecolors(blocks string, k int) int {
	ans := k
	bcnt := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'B' {
			bcnt++
		}
	}
	ans = min(ans, k-bcnt)
	for i := 1; i <= len(blocks)-k; i++ {
		if blocks[i-1] == 'B' {
			bcnt--
		}
		if blocks[i+k-1] == 'B' {
			bcnt++
		}
		ans = min(ans, k-bcnt)
	}
	return ans
}

func secondsToRemoveOccurrences(s string) int {
	ans := 0
	for {
		temp := s
		s = strings.ReplaceAll(s, "01", "10")
		if s == temp {
			return ans
		}
		ans++
	}
}

// 字母移位2
// 给定的是一段段的距离 我们想要的是每个位置的偏移量
// 区间是闭合的区间 确定下来某个位置 来统计其偏移量
// 那么这个位置实际上就是有多少段通过此 可以转化成
// 之前的位置开始的段数 减去提前结束的 使用差分 对开始进行+
// 结束进行- 反向操作反过来即可
func shiftingLetters(s string, shifts [][]int) string {
	diff := make([]int, len(s)+1)
	for _, val := range shifts {
		if val[2] == 1 {
			diff[val[0]]++
			diff[val[1]+1]--
		} else {
			diff[val[0]]--
			diff[val[1]+1]++
		}
	}
	offset := 0
	ans := make([]byte, 0)
	for i := range s {
		offset += diff[i]
		temp := byte('a' + (((int(s[i]-'a')+offset)%26)+26)%26)
		ans = append(ans, temp)
	}
	return string(ans)
}
