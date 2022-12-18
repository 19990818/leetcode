package main

import (
	"math"
)

func minMoves(nums []int, k int) int {
	// 需要列出表达式 求值
	// 根据分析实际上我们需要做的事情是将所有1的位置进行移动
	// 然后我们会移动到一个移动窗口中，然后根据距离可以知道中点会使移动距离最小
	// 距离计算公式为sum(i,k+i-1)|pj-(q+j-i)|
	// 其中pj表示为1的位置，q为移动之后序列开始的位置 我们做下变换
	// 可以得到其中的子项|pj-j-(q-i)| 其中q-i应该为pj-j的中点处的值
	// 设pj-j为gj q-i为r 那么就是一个|gj-r|的绝对值表达式 可以知道在中点取得最小值
	// 那么r=g(i+k/2)
	g, preSum := make([]int, 0), make([]int, 1)
	for i, v := range nums {
		if v == 1 {
			g = append(g, i-len(g))
			preSum = append(preSum, preSum[len(preSum)-1]+g[len(g)-1])
		}
	}
	m := len(g)
	res := math.MaxInt64
	for i := 0; i < m-k+1; i++ {
		r := g[i+k/2]
		mid := i + k/2
		f := preSum[k+i] - preSum[mid+1] + preSum[i] - preSum[mid] + r*(1-k%2)
		res = min(res, f)
	}
	return res
}

func similarPairs(words []string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isSimlar(words[i], words[j]) {
				res++
			}
		}
	}
	return res
}
func isSimlar(a, b string) bool {
	m1, m2 := make(map[rune]int), make(map[rune]int)
	for _, v := range a {
		m1[v] = 1
	}
	for _, v := range b {
		m2[v] = 1
	}
	if len(m1) != len(m2) {
		return false
	}
	for k := range m1 {
		if m2[k] == 0 {
			return false
		}
	}
	return true
}

func smallestValue(n int) int {
	primes := initPrimes()
	// 再来一个数组表示每个最小的smallestValue
	res := 0
	for n != 1 {
		temp := n
		res = 0
		for _, v := range primes {
			for n%v == 0 {
				res += v
				n = n / v
			}
			if n == 1 || v > n {
				break
			}
		}
		if res != temp {
			n = res
		}
	}
	return res
}
func initPrimes() []int {
	primes := make([]int, 0)
	for i := 2; i < 100001; i++ {
		flag := true
		for _, v := range primes {
			if v*v > i {
				break
			}
			if i%v == 0 {
				flag = false
				break
			}
		}
		if flag {
			primes = append(primes, i)
		}
	}
	return primes
}
